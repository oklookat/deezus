package deezerauth

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

var (
	// If the user clicks on "don't allow" your application is not authorized.
	ErrUserDenied = errors.New("user_denied")

	// ???.
	ErrUnknown = errors.New("unknown error")

	// Wrong response.
	ErrWrongResponse = errors.New("wrong response")

	// States not equals.
	ErrWrongState = errors.New("wrong state")
)

type Permission string

func (e Permission) String() string {
	return string(e)
}

const (
	// Access users basic information.
	//
	// Incl. name, firstname, profile picture only.
	PermissionBasicAccess Permission = "basic_access"

	// Get the user's email.
	PermissionEmail Permission = "email"

	// Access user data any time.
	PermissionOfflineAccess Permission = "offline_access"

	// Manage users' library.
	//
	// Add/rename a playlist. Add/order songs in the playlist.
	PermissionManageLibrary Permission = "manage_library"

	// Manage users' friends.
	//
	// Add/remove a following/follower.
	PermissionManageCommunity Permission = "manage_community"

	// Delete library items.
	//
	// Allow the application to delete items in the user's library.
	PermissionDeleteLibrary Permission = "delete_library"

	// Allow the application to access the user's listening history.
	PermissionListeningHistory Permission = "listening_history"
)

type AuthArgs struct {
	// Random string (CSRF protect).
	State string

	// App ID.
	AppID,

	// App secret.
	Secret,

	// Localhost uri only. Example: http://localhost
	RedirectUri string

	// RedirectUri port.
	Port int

	// Optional (default BasicAccess).
	//
	// Note: token lifetime without OfflineAccess == 3600 seconds.
	// With OfflineAccess - unlimited.
	Perms []Permission

	// User must go to this url to confirm auth.
	OnURL func(url string)
}

// 1. Create app: https://developers.deezer.com/myapps
//
// Redirect URL after authentication: http://localhost:PORT_HERE
//
// 2. Set your app data in args.
func New(ctx context.Context, args AuthArgs) (*oauth2.Token, error) {

	httpErr := make(chan error)
	code := make(chan string)

	ctxServed, cancel := context.WithCancel(ctx)
	defer func() {
		cancel()
	}()

	go serve(ctxServed, func(w http.ResponseWriter, r *http.Request) {
		var err error
		var gotCode string

		defer func() {
			if err != nil {
				w.Write([]byte(err.Error()))
				httpErr <- err
				return
			}
			w.Write([]byte("ok"))
			code <- gotCode
		}()

		if err = r.ParseForm(); err != nil {
			return
		}

		stated := r.Form.Get("state")
		if stated != args.State {
			err = ErrWrongState
			return
		}

		gotCode = r.Form.Get("code")
		if len(gotCode) > 0 {
			return
		}

		errorReason := r.Form.Get("error_reason")
		if len(errorReason) > 0 {
			switch errorReason {
			case "user_denied":
				err = ErrUserDenied
				return
			}
			err = errors.New(errorReason)
			return
		}

		err = ErrUnknown
	}, args.Port)

	var permsConv []string
	for _, p := range args.Perms {
		permsConv = append(permsConv, p.String())
	}
	goToUrl := fmt.Sprintf("https://connect.deezer.com/oauth/auth.php?app_id=%s&redirect_uri=%s:%d&perms=%s&state=%s",
		args.AppID, args.RedirectUri, args.Port, strings.Join(permsConv, ","), args.State)
	go args.OnURL(goToUrl)

	var theCode string

L:
	for {
		select {
		// Check err from http handler.
		case err := <-httpErr:
			if err != nil {
				return nil, err
			}
		// Check code from handler.
		case theCode = <-code:
			break L
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	if len(theCode) == 0 {
		return nil, ErrUnknown
	}

	tokenUrl := fmt.Sprintf("https://connect.deezer.com/oauth/access_token.php?app_id=%s&secret=%s&code=%s", args.AppID, args.Secret, theCode)

	resp, err := http.Get(tokenUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	parsed, err := url.ParseQuery(string(bodyBytes))
	if err != nil {
		return nil, ErrWrongResponse
	}

	token := parsed.Get("access_token")
	expires := parsed.Get("expires")

	if len(token) == 0 || len(expires) == 0 {
		return nil, ErrWrongResponse
	}

	expiresInt, err := strconv.ParseInt(expires, 10, 64)
	if err != nil {
		return nil, ErrWrongResponse
	}

	return &oauth2.Token{
		AccessToken: token,
		TokenType:   "Bearer",
		Expiry:      time.Unix(expiresInt, 0),
	}, err
}

func serve(ctx context.Context, what http.HandlerFunc, port int) (err error) {
	mux := http.NewServeMux()
	mux.Handle("/", what)

	portStr := ":" + strconv.Itoa(port)
	srv := &http.Server{
		Addr:    portStr,
		Handler: mux,
	}

	go func() {
		err = srv.ListenAndServe()
	}()

	if err != nil {
		return err
	}

	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	err = srv.Shutdown(ctxShutDown)

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}
