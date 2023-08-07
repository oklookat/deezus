package schema

type (
	TrackResponse struct {
		ErrorInResponse

		Track
	}

	SimpleTrack struct {
		// The track's Deezer id.
		ID ID `json:"id"`

		// true if the track is readable in the player for the current user.
		Readable bool `json:"readable"`

		// The track's fulltitle.
		Title string `json:"title"`

		// The url of the track on Deezer.
		Link string `json:"link"`

		// The track's duration in seconds.
		Duration int `json:"duration"`

		// The track's Deezer rank.
		Rank int `json:"rank"`

		// Whether the track contains explicit lyrics.
		ExplicitLyrics bool `json:"explicit_lyrics"`

		//
		Artist SimpleArtist `json:"artist"`

		// (with release date).
		Album *SimpleAlbum `json:"album"`
	}

	Track struct {
		SimpleTrack

		// The track's short title.
		TitleShort string `json:"title_short"`

		// The track version.
		TitleVersion string `json:"title_version"`

		// The track isrc.
		Isrc string `json:"isrc"`

		// The share link of the track on Deezer.
		Share *string `json:"share"`

		// The position of the track in its album.
		TrackPosition int `json:"track_position"`

		// 	The track's album's disk number.
		DiskNumber int `json:"disk_number"`

		// The track's release date.
		//
		// Nil example method: AlbumTracks (because track release date same with album).
		ReleaseDate *Time `json:"release_date"`

		// 	The url of track's preview file.
		// This file contains the first 30 seconds of the track.
		Preview string `json:"preview"`

		ExplicitContentLyrics ExplicitContent `json:"explicit_content_lyrics"`
		ExplicitContentCover  ExplicitContent `json:"explicit_content_cover"`

		// Beats per minute.
		Bpm *float64 `json:"bpm"`

		// Signal strength.
		Gain *float64 `json:"gain"`

		// List of countries where the track is available.
		AvailableCountries []string `json:"available_countries"`

		// Return a list of contributors on the track.
		Contributors []Contributor `json:"contributors"`

		// If in user library.
		TimeAdd *Time `json:"time_add"`
	}
)
