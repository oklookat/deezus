package deezus

import (
	"context"
	"errors"
	"testing"

	"github.com/oklookat/deezus/schema"
)

func TestUserMe(t *testing.T) {
	ctx := context.Background()

	cl := getClient(t)
	_, err := cl.UserMe(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserMeAlbums(t *testing.T) {
	ctx := context.Background()

	cl := getClient(t)
	res, err := cl.UserMeAlbums(ctx, 0, 10)
	if err != nil {
		t.Fatal(err)
	}
	for _, sa := range res.Data {
		println(sa.Title)
	}
}

func TestUserMeArtists(t *testing.T) {
	ctx := context.Background()

	cl := getClient(t)
	res, err := cl.UserMeArtists(ctx, 0, 10)
	if err != nil {
		t.Fatal(err)
	}
	for _, sa := range res.Data {
		println(sa.Name)
	}
}

func TestUserMePlaylsits(t *testing.T) {
	ctx := context.Background()

	cl := getClient(t)
	res, err := cl.UserMePlaylists(ctx, 0, 10)
	if err != nil {
		t.Fatal(err)
	}
	for _, sa := range res.Data {
		println(sa.Title)
	}
}

func TestUserMeTracks(t *testing.T) {
	ctx := context.Background()

	cl := getClient(t)
	res, err := cl.UserMeTracks(ctx, 0, 10)
	if err != nil {
		t.Fatal(err)
	}
	for _, sa := range res.Data {
		println(sa.Title)
	}
}

func TestAddRemoveAlbums(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	ids := _albumIds[:4]

	_, err := cl.AddAlbums(ctx, ids)
	if err != nil {
		t.Fatal(err)
	}

	for _, id := range ids {
		_, err = cl.RemoveAlbum(ctx, id)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestAddRemoveArtists(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	ids := _artistIds[:4]

	_, err := cl.AddArtists(ctx, ids)
	if err != nil {
		t.Fatal(err)
	}

	for _, id := range ids {
		_, err = cl.RemoveArtist(ctx, id)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestAddRemoveTracks(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	ids := _trackIds[:4]

	_, err := cl.AddTracks(ctx, ids)
	if err != nil {
		t.Fatal(err)
	}

	for _, id := range ids {
		_, err = cl.RemoveTrack(ctx, id)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestLikeUnlikePlaylist(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	ids := _playlistIds[:4]
	if _, err := cl.LikePlaylists(ctx, ids); err != nil {
		t.Fatal(err)
	}
	for _, id := range ids {
		if _, err := cl.UnlikePlaylist(ctx, id); err != nil {
			t.Fatal(err)
		}
	}
}

func TestLikePlaylistGoodBad(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	//TestLikeUnlikePlaylist(t)
	_, err := cl.likePlaylistsBad(ctx, nil)
	if err == nil {
		t.Fatal("expected error")
	}
	if !errors.As(err, &schema.Error{}) {
		t.Fatal("expected schema.Error")
	}
}
