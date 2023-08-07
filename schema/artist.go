package schema

type (
	ArtistResponse struct {
		ErrorInResponse

		Artist
	}

	SimpleArtist struct {
		ID ID `json:"id"`

		// The artist's name.
		Name string `json:"name"`

		Picture       *string `json:"picture"`
		PictureSmall  *string `json:"picture_small"`
		PictureMedium *string `json:"picture_medium"`
		PictureBig    *string `json:"picture_big"`
		PictureXl     *string `json:"picture_xl"`
	}

	Artist struct {
		SimpleArtist

		// The url of the artist on Deezer.
		Link string `json:"link"`

		// The url of the artist picture.
		Share string `json:"share"`

		// The number of artist's albums.
		NbAlbum int `json:"nb_album"`

		// The number of artist's fans.
		NbFan int `json:"nb_fan"`

		// true if the artist has a smartradio.
		Radio bool `json:"radio"`

		// API Link to the top of this artist.
		Tracklist string `json:"tracklist"`

		// If in user library.
		TimeAdd *Time `json:"time_add"`
	}

	Contributor struct {
		SimpleArtist

		Link  string `json:"link"`
		Share string `json:"share"`
		Radio bool   `json:"radio"`

		// Example: "Main".
		Role string `json:"role"`
	}
)
