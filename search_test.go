package deezus

import (
	"context"
	"testing"
)

func TestSearchAlbums(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.SearchAlbums(ctx, "Тальник Cyou", "", false, 0, 5)
	if err != nil {
		t.Fatal(err)
	}
	for _, sa := range resp.Data {
		println(sa.Title)
	}
}

func TestSearchArtists(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.SearchArtists(ctx, "eminem", "", false, 0, 5)
	if err != nil {
		t.Fatal(err)
	}
	for _, sa := range resp.Data {
		println(sa.Name)
	}
}

func TestSearchPlaylists(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.SearchPlaylists(ctx, "rap", "", false, 0, 5)
	if err != nil {
		t.Fatal(err)
	}
	for _, sa := range resp.Data {
		println(sa.Title)
	}
}

func TestSearchTracks(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.SearchTracks(ctx, "eminem", "", false, 0, 5)
	if err != nil {
		t.Fatal(err)
	}
	for _, sa := range resp.Data {
		println(sa.Title)
	}
}
