package deezus

import (
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/oklookat/deezus/schema"
)

var (
	_artistIds = [4]schema.ID{
		351586,
		106110,
		5254467,
		1484833,
	}
	_albumIds = [4]schema.ID{
		282741712,
		266806782,
		298052682,
		57475222,
	}
	_trackIds = [4]schema.ID{
		464345862,
		138546553,
		3247949,
		4286051,
	}
	_playlistIds = [4]schema.ID{
		8190352822,
		9341070582,
		8971689702,
		8584980222,
	}
)

func getClient(t *testing.T) *Client {
	err := godotenv.Load()
	if err != nil {
		t.Fatal(err)
	}

	cl, err := New(os.Getenv("TOKEN"))
	if err != nil {
		t.Fatal(err)
	}
	cl.Http.SetRateLimit(5, time.Second)
	//cl.Http.SetLogger(loggerDefault{})

	return cl
}

// type loggerDefault struct {
// }

// func (l loggerDefault) Debugf(msg string, args ...any) {
// 	log.Printf(msg, args...)
// }

// func (l loggerDefault) Err(msg string, err error) {
// 	if err == nil {
// 		log.Printf("%s", msg)
// 		return
// 	}
// 	log.Printf("%s. Err: %s", msg, err.Error())
// }
