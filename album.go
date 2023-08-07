package deezus

import (
	"context"
	"net/http"

	"github.com/oklookat/deezus/schema"
)

// https://developers.deezer.com/api/album#infos
// https://developers.deezer.com/api/album#connections

// Get album by id.
func (c *Client) Album(ctx context.Context, id schema.ID) (*schema.AlbumResponse, error) {
	data := &schema.AlbumResponse{}
	resp, err := httpAny(ctx, c, nil, data, nil, http.MethodGet, "album", id.String())
	if err == nil {
		err = checkResponse(resp, data.ErrorInResponse)
	}
	return data, err
}

// Return a list of album's tracks. Represented by an array of Track objects.
func (c Client) AlbumTracks(ctx context.Context, id schema.ID, index, limit int) (*schema.Response[[]schema.Track], error) {
	return getAnyResp[[]schema.Track](ctx, &c, getIndexLimit(index, limit), "album", id.String(), "tracks")
}
