package dto

import "github.com/ruziba3vich/sahiy_management/internal/domain/user"

type LoginRequest struct {
	Phone    string `json:"phone" binding:"required" example:"+998901112233"`
	Password string `json:"password" binding:"required" example:"SuperSecure2026!"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type RefreshTokenRequest struct {
	Token string `json:"token" binding:"required"`
}

type RefreshTokenResponse struct {
	Token string `json:"token"`
}

func ToLoginResponse(token string, u *user.User) *LoginResponse {
	return &LoginResponse{
		Token: token,
		User:  *ToUserResponse(u),
	}
}
