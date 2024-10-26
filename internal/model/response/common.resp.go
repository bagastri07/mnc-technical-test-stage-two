package response

type MessageResp struct {
	Message string `json:"message"`
}

type TokenResponse struct {
	Token     string `json:"token"`
	ExpiredAt string `json:"expired_at"`
}
