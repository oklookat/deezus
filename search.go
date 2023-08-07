package deezus

import (
	"context"

	"github.com/oklookat/deezus/schema"
)

// https://developers.deezer.com/api/search
// https://developers.deezer.com/api/search#connections

// Search albums. Order can be empty.
func (c Client) SearchAlbums(ctx context.Context, query string, order schema.Order, strict bool, index, limit int) (*schema.Response[[]schema.SimpleAlbum], error) {
	params := getIndexLimit(index, limit)
	setSearchParams(params, query, order, strict)
	return getAnyResp[[]schema.SimpleAlbum](ctx, &c, params, "search", "album")
}

// Search artists. Order can be empty.
func (c Client) SearchArtists(ctx context.Context, query string, order schema.Order, strict bool, index, limit int) (*schema.Response[[]schema.SimpleArtist], error) {
	params := getIndexLimit(index, limit)
	setSearchParams(params, query, order, strict)
	return getAnyResp[[]schema.SimpleArtist](ctx, &c, params, "search", "artist")
}

// Search playlists. Order can be empty.
func (c Client) SearchPlaylists(ctx context.Context, query string, order schema.Order, strict bool, index, limit int) (*schema.Response[[]schema.SimplePlaylist], error) {
	params := getIndexLimit(index, limit)
	setSearchParams(params, query, order, strict)
	return getAnyResp[[]schema.SimplePlaylist](ctx, &c, params, "search", "playlist")
}

// Search tracks. Order can be empty.
func (c Client) SearchTracks(ctx context.Context, query string, order schema.Order, strict bool, index, limit int) (*schema.Response[[]schema.SimpleTrack], error) {
	params := getIndexLimit(index, limit)
	setSearchParams(params, query, order, strict)
	return getAnyResp[[]schema.SimpleTrack](ctx, &c, params, "search", "track")
}
