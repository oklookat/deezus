package deezus

import (
	"context"
	"testing"
)

func TestArtist(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.Artist(ctx, _artistIds[0])
	if err != nil {
		t.Fatal(err)
	}
	println(resp.Name)
}

func TestArtistTop(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.ArtistTop(ctx, _artistIds[0], 0, 10)
	if err != nil {
		t.Fatal(err)
	}
	for _, t2 := range resp.Data {
		println(t2.Title)
	}
}

func TestArtistAlbums(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.ArtistAlbums(ctx, _artistIds[0], 0, 10)
	if err != nil {
		t.Fatal(err)
	}
	for _, t2 := range resp.Data {
		println(t2.Title)
	}
}
