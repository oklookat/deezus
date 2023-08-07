package deezus

import (
	"context"
	"net/url"
	"time"

	"github.com/oklookat/deezus/schema"
	"github.com/oklookat/vantuz"
)

// Query quota: 50 requests / 5 seconds.
func New(accessToken string) (*Client, error) {
	vals := url.Values{}
	vals.Set("access_token", accessToken)
	httpCl := vantuz.C().SetGlobalQueryParams(vals)
	httpCl.SetRateLimit(50, 5*time.Second)

	cl := &Client{
		Http: httpCl,
	}
	resp, err := cl.UserMe(context.Background())
	if err != nil {
		return nil, err
	}
	cl.UserID = resp.ID

	return cl, err
}

type Client struct {
	// Current user ID.
	UserID schema.ID

	Http *vantuz.Client
}

func (c *Client) SetUserAgent(val string) {
	c.Http.SetUserAgent(val)
}
