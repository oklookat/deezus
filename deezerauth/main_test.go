package auth

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

func TestNew(t *testing.T) {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		t.Fatal(err)
	}

	appId := os.Getenv("APP_ID")
	secret := os.Getenv("SECRET")
	redirectUri := os.Getenv("REDIRECT_URI")
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		t.Fatal(err)
	}

	perms := []Permission{PermissionBasicAccess,
		PermissionEmail,
		PermissionOfflineAccess,
		PermissionManageLibrary,
		PermissionManageCommunity,
		PermissionDeleteLibrary,
		PermissionListeningHistory,
	}
	args := AuthArgs{
		State:       "123",
		AppID:       appId,
		Secret:      secret,
		RedirectUri: redirectUri,
		Port:        port,
		Perms:       perms,
		OnURL: func(url string) {
			println(url)
		},
	}

	tok, err := New(ctx, args)
	if err != nil {
		t.Fatal(err)
	}
	println(tok.TokenType)
}
