package schema

type (
	UserMeResponse struct {
		ErrorInResponse

		UserMe
	}

	UserMe struct {
		ID                             ID       `json:"id"`
		Name                           string   `json:"name"`
		Lastname                       string   `json:"lastname"`
		Firstname                      string   `json:"firstname"`
		Email                          string   `json:"email"`
		Status                         int      `json:"status"`
		Birthday                       string   `json:"birthday"`
		InscriptionDate                string   `json:"inscription_date"`
		Gender                         string   `json:"gender"`
		Link                           string   `json:"link"`
		Picture                        string   `json:"picture"`
		PictureSmall                   string   `json:"picture_small"`
		PictureMedium                  string   `json:"picture_medium"`
		PictureBig                     string   `json:"picture_big"`
		PictureXl                      string   `json:"picture_xl"`
		Country                        string   `json:"country"`
		Lang                           string   `json:"lang"`
		IsKid                          bool     `json:"is_kid"`
		ExplicitContentLevel           string   `json:"explicit_content_level"`
		ExplicitContentLevelsAvailable []string `json:"explicit_content_levels_available"`
		Tracklist                      string   `json:"tracklist"`
	}
)
