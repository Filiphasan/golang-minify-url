package request

import validation "github.com/go-ozzo/ozzo-validation"

type GetShortenedURLRequest struct {
	Token string `json:"token"`
}

func (r *GetShortenedURLRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Token,
			validation.Required.Error("Token is required!"),
			validation.Length(7, 15).Error("Token length must be between 7 and 15 characters!"),
		),
	)
}
