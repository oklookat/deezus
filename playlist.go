package deezus

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/oklookat/deezus/schema"
)

// https://developers.deezer.com/api/playlist
// https://developers.deezer.com/api/playlist#actions
// https://developers.deezer.com/api/playlist#connections

// Get playlist by ID.
func (c *Client) Playlist(ctx context.Context, id schema.ID) (*schema.PlaylistResponse, error) {
	data := &schema.PlaylistResponse{}
	resp, err := httpAny(ctx, c, nil, data, nil, http.MethodGet, "playlist", id.String())
	if err == nil {
		err = checkResponse(resp, data.ErrorInResponse)
	}
	return data, err
}

// Get playlist tracks.
func (c Client) PlaylistTracks(ctx context.Context, id schema.ID, index, limit int) (*schema.Response[[]schema.SimpleTrack], error) {
	return getAnyResp[[]schema.SimpleTrack](ctx, &c, getIndexLimit(index, limit), "playlist", id.String(), "tracks")
}

// Add a track(s) to the playlist.
func (c *Client) AddTracksToPlaylist(ctx context.Context, id schema.ID, tracks []schema.ID) (*schema.BoolResponse, error) {
	body := url.Values{}
	body.Set("songs", idsJoin(tracks))
	return postAny(ctx, c, body, "playlist", id.String(), "tracks")
}

// Order tracks in the playlist.
func (c *Client) OrderTracksInPlaylist(ctx context.Context, id schema.ID, tracks []schema.ID) (*schema.BoolResponse, error) {
	body := url.Values{}
	body.Set("order", idsJoin(tracks))
	return postAny(ctx, c, body, "playlist", id.String(), "tracks")
}

// Remove track(s) from the playlist.
func (c *Client) RemoveTracksFromPlaylist(ctx context.Context, id schema.ID, tracks []schema.ID) (*schema.BoolResponse, error) {
	body := url.Values{}
	body.Set("songs", idsJoin(tracks))
	return deleteAny(ctx, c, body, "playlist", id.String(), "tracks")
}

// Update playlist. Nil title; desc; bool will be ignored.
func (c *Client) UpdatePlaylist(ctx context.Context, id schema.ID, title, description *string, isPublic *bool) (*schema.BoolResponse, error) {
	if title == nil && description == nil && isPublic == nil {
		return nil, nil
	}
	params := url.Values{}
	if title != nil && len(*title) > 0 {
		params.Set("title", *title)
	}
	if description != nil {
		params.Set("description", *description)
	}
	if isPublic != nil {
		params.Set("public", strconv.FormatBool(*isPublic))
	}
	return postAnyParams(ctx, c, params, "playlist", id.String())
}

// Delete the playlist.
func (c *Client) DeletePlaylist(ctx context.Context, id schema.ID) (*schema.BoolResponse, error) {
	return deleteAny(ctx, c, nil, "playlist", id.String())
}
