package request

import validation "github.com/go-ozzo/ozzo-validation"

type SetShortenURLRequest struct {
	Url       string `json:"url"`
	ExpireDay int    `json:"expireDay"`
	HasQrCode bool   `json:"hasQrCode"`
}

func (r *SetShortenURLRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(r.Url,
			validation.Required.Error("Url is required"),
			validation.Length(1, 1000).Error("Url length must be between 1 and 1000 characters"),
		),
		validation.Field(r.ExpireDay,
			validation.Max(365).Error("Expire day must be less than or equal to 365"),
		),
	)
}
