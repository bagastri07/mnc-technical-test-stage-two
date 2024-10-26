package request

type UserLoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegisterReq struct {
	Username string  `json:"username" binding:"required"`
	Password string  `json:"password" binding:"required"`
	FullName *string `json:"full_name"`
}

type GetUserInfoReq struct {
	Username string
}
