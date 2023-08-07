package deezus

import (
	"context"
	"net/http"

	"github.com/oklookat/deezus/schema"
)

// https://developers.deezer.com/api/artist
// https://developers.deezer.com/api/artist#connections

// Get artist by id.
func (c *Client) Artist(ctx context.Context, id schema.ID) (*schema.ArtistResponse, error) {
	data := &schema.ArtistResponse{}
	resp, err := httpAny(ctx, c, nil, data, nil, http.MethodGet, "artist", id.String())
	if err == nil {
		err = checkResponse(resp, data.ErrorInResponse)
	}
	return data, err
}

// Get tracks of an artist (starting from top tracks).
func (c Client) ArtistTop(ctx context.Context, id schema.ID, index, limit int) (*schema.Response[[]schema.SimpleTrack], error) {
	return getAnyResp[[]schema.SimpleTrack](ctx, &c, getIndexLimit(index, limit), "artist", id.String(), "top")
}

// Return a list of artist's albums. Represented by an array of Album objects.
func (c Client) ArtistAlbums(ctx context.Context, id schema.ID, index, limit int) (*schema.Response[[]schema.SimpleAlbum], error) {
	return getAnyResp[[]schema.SimpleAlbum](ctx, &c, getIndexLimit(index, limit), "artist", id.String(), "albums")
}
