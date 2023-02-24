package schema

const TokenType = "Bearer"

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResUser struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
}

type LoginRes struct {
	AccessToken string       `json:"access_token"`
	TokenType   string       `json:"token_type"`
	User        LoginResUser `json:"user"`
}

type ResetPasswordReq struct {
	Email string `json:"email"`
}

type UpdatePasswordReq struct {
	Password string `json:"password"`
}
