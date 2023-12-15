package mapper

import (
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/domain"
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/model"
)

func ToUserResponse(user *domain.User) *model.UserResponse {
	return &model.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Username: user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}