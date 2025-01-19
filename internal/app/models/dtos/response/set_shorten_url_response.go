package response

type SetShortenURLResponse struct {
	Token        string `json:"token"`
	ShortenedUrl string `json:"shortenedUrl"`
	QrCode       string `json:"qrCode"`
}
