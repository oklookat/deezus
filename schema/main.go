package schema

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// The Deezer id (string).
type ID int64

func (e ID) String() string {
	return strconv.FormatInt(int64(e), 10)
}

func (e ID) Int64() int64 {
	return int64(e)
}

const (
	ApiUrl = "https://api.deezer.com"
)

// The explicit content values.
type ExplicitContent int

const (
	ExplicitContentNotExplicit ExplicitContent = 0
	ExplicitContentExplicit    ExplicitContent = 1
	ExplicitContentUnknown     ExplicitContent = 2
	ExplicitContentEdited      ExplicitContent = 3

	// (Album "lyrics" only).
	ExplicitContentPartiallyExplicit ExplicitContent = 4

	// (Album "lyrics" only).
	ExplicitContentPartiallyUnknown ExplicitContent = 5

	ExplicitContentNoAdviceAvailable ExplicitContent = 6

	// (Album "lyrics" only).
	ExplicitContentPartiallyNoAdviceAvailable ExplicitContent = 7
)

type RecordType string

const (
	RecordTypeAlbum = "album"
	RecordTypeEP    = "ep"
)

func (e RecordType) String() string {
	return string(e)
}

type PictureSize string

func (e PictureSize) String() string {
	return string(e)
}

const (
	PictureSizeSmall  PictureSize = "small"
	PictureSizeMedium PictureSize = "medium"
	PictureSizeBig    PictureSize = "big"
	PictureSizeXl     PictureSize = "xl"
)

func GetPictureURL(picUrl string, size PictureSize) *url.URL {
	sized, _ := url.Parse(picUrl)
	if sized == nil {
		return nil
	}
	q := sized.Query()
	q.Add("size", size.String())
	sized.RawQuery = q.Encode()
	return sized
}

type (
	ErrorInResponse struct {
		Error *Error `json:"error"`
	}

	IDResponse struct {
		ErrorInResponse

		ID ID `json:"id"`
	}

	Response[T any] struct {
		ErrorInResponse

		Data T `json:"data"`

		// Example: available when search.
		Total *int `json:"total"`

		// Example: available when search.
		Next *string `json:"next"`
	}
)

type Error struct {
	Type    string    `json:"type"`
	Message string    `json:"message"`
	Code    ErrorCode `json:"code"`
}

func (e Error) Error() string {
	return e.Message
}

type ErrorCode int

const (
	ErrorCodeQuota              ErrorCode = 4
	ErrorCodeItemsLimitExceeded ErrorCode = 100
	ErrorCodePermission         ErrorCode = 200
	ErrorCodeTokenInvalid       ErrorCode = 300
	ErrorCodeParameter          ErrorCode = 500
	ErrorCodeParameterMissing   ErrorCode = 501
	ErrorCodeQueryInvalid       ErrorCode = 600
	ErrorCodeServiceBusy        ErrorCode = 700
	// Example: artist not found; wrong api path; etc.
	ErrorCodeDataNotFound                ErrorCode = 800
	ErrorCodeIndividualAccountNotAllowed ErrorCode = 901
)

type Time time.Time

func (e Time) Time() time.Time {
	return time.Time(e)
}

func (e *Time) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	isTimeOk := func(t time.Time, err error) bool {
		return err == nil && t.Year() <= time.Now().Year()
	}

	convTime := func(t time.Time) {
		conv := Time(t)
		*e = conv
	}

	// String.
	var str string
	if err := json.Unmarshal(b, &str); err == nil {
		// 2022-02-13
		t, err := time.Parse(time.DateOnly, str)
		if isTimeOk(t, err) {
			convTime(t)
			return nil
		}

		// 2022-02-13 02:34:26
		t, err = time.Parse(time.DateTime, str)
		if isTimeOk(t, err) {
			convTime(t)
			return nil
		}
	}

	// Unix.
	var i64 int64
	if err := json.Unmarshal(b, &i64); err == nil {
		t := time.Unix(i64, 0)
		if isTimeOk(t, err) {
			convTime(t)
			return nil
		}
	}

	return errors.New("failed to parse schema.Time")
}

type Order string

func (e Order) String() string {
	return string(e)
}

const (
	OrderRanking      Order = "RANKING"
	OrderTrackAsc     Order = "TRACK_ASC"
	OrderTrackDesc    Order = "TRACK_DESC"
	OrderArtistAsc    Order = "ARTIST_ASC"
	OrderArtistDesc   Order = "ARTIST_DESC"
	OrderAlbumAsc     Order = "ALBUM_ASC"
	OrderAlbumDesc    Order = "ALBUM_DESC"
	OrderRatingAsc    Order = "RATING_ASC"
	OrderRatingDesc   Order = "RATING_DESC"
	OrderDurationAsc  Order = "DURATION_ASC"
	OrderDurationDesc Order = "DURATION_DESC"
)

type BoolResponse struct {
	ErrorInResponse

	Result *bool
}

func (e *BoolResponse) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	boold, err := strconv.ParseBool(string(b))
	if err == nil {
		e.Result = &boold
		return nil
	}
	respErr := &Response[any]{}
	if err = json.Unmarshal(b, respErr); err == nil {
		e.ErrorInResponse.Error = respErr.Error
		return nil
	}
	return fmt.Errorf("boolResponse.UnmarshalJSON: %w", err)
}
