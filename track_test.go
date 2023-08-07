package deezus

import (
	"context"
	"testing"
)

func TestTrack(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	resp, err := cl.Track(ctx, _trackIds[0])
	if err != nil {
		t.Fatal(err)
	}
	println(resp.Title)
}
