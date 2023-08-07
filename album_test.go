package deezus

import (
	"context"
	"testing"
)

func TestAlbum(t *testing.T) {
	ctx := context.Background()

	cl := getClient(t)
	res, err := cl.Album(ctx, _albumIds[0])
	if err != nil {
		t.Fatal(err)
	}
	println(res.Title)
}

func TestAlbumTracks(t *testing.T) {
	ctx := context.Background()

	cl := getClient(t)
	res, err := cl.AlbumTracks(ctx, _albumIds[0], 0, 10)
	if err != nil {
		t.Fatal(err)
	}
	for _, sa := range res.Data {
		println(sa.Title)
	}
}
