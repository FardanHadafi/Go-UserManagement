package service

import (
	"Go-UserManagement/model/web"
	"context"
)

// Same as Repository
type UserService interface {
	Register(ctx context.Context, r web.UserRegisterRequest) (web.UserResponse, error)
	Login(ctx context.Context, request web.UserLoginRequest) (web.UserResponse, error)
}
