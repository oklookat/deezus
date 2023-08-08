package schema

type (
	AlbumResponse struct {
		ErrorInResponse

		Album
	}

	SimpleAlbum struct {
		// The Deezer album id.
		ID ID `json:"id"`

		// The album title.
		Title string `json:"title"`

		// The url of the album's cover. Add 'size' parameter to the url to change size.
		// Can be 'small', 'medium', 'big', 'xl'
		Cover string `json:"cover"`

		// The url of the album's cover in size small.
		CoverSmall string `json:"cover_small"`

		// The url of the album's cover in size medium.
		CoverMedium string `json:"cover_medium"`

		// The url of the album's cover in size big.
		CoverBig string `json:"cover_big"`

		// The url of the album's cover in size xl.
		CoverXl string `json:"cover_xl"`

		// The record type of the album.
		//
		// Not nil examples: ArtistAlbums.
		RecordType *RecordType `json:"record_type"`

		// The album's release date.
		//
		// Not nil examples: ArtistAlbums.
		ReleaseDate *Time `json:"release_date"`
	}

	Album struct {
		SimpleAlbum

		// The album's release date.
		ReleaseDate *Time `json:"release_date"`

		// The album UPC.
		Upc string `json:"upc"`

		// The url of the album on Deezer.
		Link string `json:"link"`

		// The share link of the album on Deezer.
		Share string `json:"share"`

		//
		Md5Image string `json:"md5_image"`

		// The album's first genre id (You should use the genre list instead).
		// NB : -1 for not found.
		GenreID GenreID `json:"genre_id"`

		// List of genre object.
		Genres struct {
			Data []Genre `json:"data"`
		} `json:"genres"`

		// The album's label name.
		Label string `json:"label"`

		//
		NbTracks int `json:"nb_tracks"`

		// The album's duration (seconds).
		Duration int `json:"duration"`

		// The number of album's Fans.
		Fans int `json:"fans"`

		// Return an alternative album object if the current album is not available.
		Available bool `json:"available"`

		// API Link to the tracklist of this album.
		Tracklist string `json:"tracklist"`

		// Whether the album contains explicit lyrics.
		ExplicitLyrics bool `json:"explicit_lyrics"`

		// The explicit content lyrics values.
		ExplicitContentLyrics ExplicitContent `json:"explicit_content_lyrics"`

		// The explicit cover values.
		ExplicitContentCover ExplicitContent `json:"explicit_content_cover"`

		// Return a list of contributors on the album.
		Contributors []Contributor `json:"contributors"`

		Artist SimpleArtist `json:"artist"`

		// If in user library.
		TimeAdd *Time `json:"time_add"`

		//
		Tracks struct {
			Data []SimpleTrack `json:"data"`
		} `json:"tracks"`
	}
)
