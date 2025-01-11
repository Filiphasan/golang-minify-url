package response

type HealthRes struct {
	Healthy bool          `json:"healthy"`
	Redis   HealthResItem `json:"redis"`
	Mongo   HealthResItem `json:"mongo"`
}

type HealthResItem struct {
	Message string `json:"message"`
	Healthy bool   `json:"healthy"`
}
