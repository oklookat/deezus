package schema

type (
	PlaylistResponse struct {
		ErrorInResponse

		Playlist
	}

	SimplePlaylist struct {
		ID     ID     `json:"id"`
		Title  string `json:"title"`
		Public bool   `json:"public"`
		// Liked tracks playlist?
		IsLovedTrack  bool   `json:"is_loved_track"`
		Collaborative bool   `json:"collaborative"`
		NbTracks      int    `json:"nb_tracks"`
		Picture       string `json:"picture"`
		PictureSmall  string `json:"picture_small"`
		PictureMedium string `json:"picture_medium"`
		PictureBig    string `json:"picture_big"`
		PictureXl     string `json:"picture_xl"`

		// Not nil example: UserMePlaylists.
		Creator *struct {
			ID   ID     `json:"id"`
			Name string `json:"name"`
		} `json:"creator"`
	}

	Playlist struct {
		SimplePlaylist

		Description  string `json:"description"`
		Duration     int    `json:"duration"`
		Fans         int    `json:"fans"`
		Link         string `json:"link"`
		Share        string `json:"share"`
		Checksum     string `json:"checksum"`
		Tracklist    string `json:"tracklist"`
		CreationDate *Time  `json:"creation_date"`
		Tracks       struct {
			Data     []SimpleTrack `json:"data"`
			Checksum string        `json:"checksum"`
		} `json:"tracks"`
	}
)
