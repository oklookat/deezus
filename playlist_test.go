package deezus

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/oklookat/deezus/schema"
)

func TestPlaylist(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	res, err := cl.Playlist(ctx, _playlistIds[0])
	if err != nil {
		t.Fatal(err)
	}
	println(res.Title)
}

func TestPlaylistTracks(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.PlaylistTracks(ctx, _playlistIds[0], 0, 5)
	if err != nil {
		t.Fatal(err)
	}
	for _, t2 := range resp.Data {
		println(t2.Title)
	}
}

func TestPlaylistCRUD(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	// Create.
	ideed, err := cl.CreatePlaylist(ctx, "test")
	if err != nil {
		t.Fatal(err)
	}

	// Read.
	if _, err = cl.Playlist(ctx, ideed.ID); err != nil {
		t.Fatal(err)
	}

	// Add tracks.
	trackIds := _trackIds[:4]
	if _, err = cl.AddTracksToPlaylist(ctx, ideed.ID, _trackIds[:4]); err != nil {
		t.Fatal(err)
	}

	// Order tracks.
	trackIdsShuffled := make([]schema.ID, len(trackIds))
	copy(trackIdsShuffled, trackIds)
	rand.New(rand.NewSource(time.Now().Unix()))
	rand.Shuffle(len(trackIdsShuffled), func(i, j int) {
		trackIdsShuffled[i], trackIdsShuffled[j] = trackIdsShuffled[j], trackIdsShuffled[i]
	})
	if _, err = cl.OrderTracksInPlaylist(ctx, ideed.ID, trackIdsShuffled); err != nil {
		t.Fatal(err)
	}

	// Remove tracks.
	if _, err = cl.RemoveTracksFromPlaylist(ctx, ideed.ID, trackIds); err != nil {
		t.Fatal(err)
	}

	// Update playlist.
	title := "test2"
	desc := "cooldesc"
	isPub := true
	if _, err = cl.UpdatePlaylist(ctx, ideed.ID, &title, &desc, &isPub); err != nil {
		t.Fatal(err)
	}

	// Delete playlist.
	if _, err = cl.DeletePlaylist(ctx, ideed.ID); err != nil {
		t.Fatal(err)
	}
}
