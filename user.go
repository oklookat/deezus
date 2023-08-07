package deezus

import (
	"context"
	"net/http"
	"net/url"

	"github.com/oklookat/deezus/schema"
)

// https://developers.deezer.com/api/user
// https://developers.deezer.com/api/user#connections

// Get current user.
func (c *Client) UserMe(ctx context.Context) (*schema.UserMeResponse, error) {
	resp := &schema.UserMeResponse{}
	hResp, err := httpAny(ctx, c, nil, resp, nil, http.MethodGet, "user", "me")
	if err != nil {
		return nil, err
	}
	return resp, checkResponse(hResp, resp.ErrorInResponse)
}

// Return a list of user's favorite albums. Represented by an array of Album object.
func (c Client) UserMeAlbums(ctx context.Context, index, limit int) (*schema.Response[[]schema.SimpleAlbum], error) {
	return getAnyResp[[]schema.SimpleAlbum](ctx, &c, getIndexLimit(index, limit), "user", "me", "albums")
}

// Return a list of user's favorite artists. Represented by an array of Artist object.
func (c Client) UserMeArtists(ctx context.Context, index, limit int) (*schema.Response[[]schema.SimpleArtist], error) {
	return getAnyResp[[]schema.SimpleArtist](ctx, &c, getIndexLimit(index, limit), "user", "me", "artists")
}

// Return a list of user's public Playlist, represented by an array of Playlist object.
// Permission is needed to return private playlists
func (c Client) UserMePlaylists(ctx context.Context, index, limit int) (*schema.Response[[]schema.SimplePlaylist], error) {
	return getAnyResp[[]schema.SimplePlaylist](ctx, &c, getIndexLimit(index, limit), "user", "me", "playlists")
}

// Return a list of user's favorite tracks. Represented by an array of Track object.
func (c Client) UserMeTracks(ctx context.Context, index, limit int) (*schema.Response[[]schema.SimpleTrack], error) {
	return getAnyResp[[]schema.SimpleTrack](ctx, &c, getIndexLimit(index, limit), "user", "me", "tracks")
}

// Add album(s) to the user's library.
//
// If album(s) not found, no errors appear.
func (c *Client) AddAlbums(ctx context.Context, ids []schema.ID) (*schema.BoolResponse, error) {
	body := url.Values{}
	body.Set("album_ids", idsJoin(ids))
	return postAny(ctx, c, body, "user", "me", "albums")
}

// Remove album from user library.
//
// If album not found, no errors appear.
func (c *Client) RemoveAlbum(ctx context.Context, id schema.ID) (*schema.BoolResponse, error) {
	params := url.Values{}
	params.Set("album_id", id.String())
	return deleteAny(ctx, c, params, "user", "me", "albums")
}

// Add artist(s) to the user's library.
//
// If artist(s) not found, no errors appear.
func (c *Client) AddArtists(ctx context.Context, ids []schema.ID) (*schema.BoolResponse, error) {
	body := url.Values{}
	body.Set("artist_ids", idsJoin(ids))
	return postAny(ctx, c, body, "user", "me", "artists")
}

// Remove album from user library.
//
// If album not found, no errors appear.
func (c *Client) RemoveArtist(ctx context.Context, id schema.ID) (*schema.BoolResponse, error) {
	params := url.Values{}
	params.Set("artist_id", id.String())
	return deleteAny(ctx, c, params, "user", "me", "artists")
}

// Add track(s) to the user's library.
//
// If track(s) not found, no errors appear.
func (c *Client) AddTracks(ctx context.Context, ids []schema.ID) (*schema.BoolResponse, error) {
	body := url.Values{}
	body.Set("track_ids", idsJoin(ids))
	return postAny(ctx, c, body, "user", "me", "tracks")
}

// Remove album from user library.
//
// If album not found, no errors appear.
func (c *Client) RemoveTrack(ctx context.Context, id schema.ID) (*schema.BoolResponse, error) {
	params := url.Values{}
	params.Set("track_id", id.String())
	return deleteAny(ctx, c, params, "user", "me", "tracks")
}

// Create a playlist.
func (c *Client) CreatePlaylist(ctx context.Context, title string) (*schema.IDResponse, error) {
	params := url.Values{}
	params.Set("title", title)
	data := &schema.IDResponse{}
	resp, err := httpAny(ctx, c, params, data, nil, http.MethodPost, "user", "me", "playlists")
	if err == nil {
		err = checkResponse(resp, data.ErrorInResponse)
	}
	return data, err
}

// Add playlist(s) to the user's favorites.
func (c *Client) LikePlaylists(ctx context.Context, ids []schema.ID) (*schema.BoolResponse, error) {
	params := url.Values{}
	params.Set("playlist_ids", idsJoin(ids))
	return postAnyParams(ctx, c, params, "user", "me", "playlists")
}

// Debug (schema.BoolResponse testing).
func (c *Client) likePlaylistsBad(ctx context.Context, ids []schema.ID) (*schema.BoolResponse, error) {
	params := url.Values{}
	//params.Set("playlist_ids", idsJoin(ids))
	return postAnyParams(ctx, c, params, "user", "me", "playlists")
}

// Remove a playlist from the user's favorites.
func (c *Client) UnlikePlaylist(ctx context.Context, id schema.ID) (*schema.BoolResponse, error) {
	params := url.Values{}
	params.Set("playlist_id", id.String())
	return deleteAny(ctx, c, params, "user", "me", "playlists")
}
