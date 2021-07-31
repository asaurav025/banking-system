package dto

type JwtCreationResponse struct {
	Token    string `json:"token"`
	Duration string `json:"duration"`
}
