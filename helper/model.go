package helper

import (
	"Go-UserManagement/model/domain"
	"Go-UserManagement/model/web"
)

func ToUserResponse(User domain.User) web.UserResponse {
	return web.UserResponse{
		ID: User.ID,
		Name: User.Name,
	}
}