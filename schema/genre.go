package schema

type (
	GenreID int

	Genre struct {
		ID      GenreID `json:"id"`
		Name    string  `json:"name"`
		Picture string  `json:"picture"`
		Type    string  `json:"type"`
	}
)
