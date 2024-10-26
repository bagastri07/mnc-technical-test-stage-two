package response

type MessageResp struct {
	Message string `json:"message"`
}

type PostLoginUserResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type CommonResultResp struct {
	Status string `json:"status"`
	Result any    `json:"result"`
}
