package deezus

import (
	"context"
	"net/http"

	"github.com/oklookat/deezus/schema"
)

// https://developers.deezer.com/api/track
// https://developers.deezer.com/api/track#actions

// Get track by id.
func (c *Client) Track(ctx context.Context, id schema.ID) (*schema.TrackResponse, error) {
	data := &schema.TrackResponse{}
	resp, err := httpAny(ctx, c, nil, data, nil, http.MethodGet, "track", id.String())
	if err == nil {
		err = checkResponse(resp, data.ErrorInResponse)
	}
	return data, err
}
