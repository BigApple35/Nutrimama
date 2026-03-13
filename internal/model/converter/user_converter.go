package converter

import (
	"nutrimama/internal/entity"
	"nutrimama/internal/model"
)

func UserToResponse(user *entity.User, token string) *model.UserResponse {
	return &model.UserResponse{
		UserId:      user.ID,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
		Token:     	 token,
	}
}
