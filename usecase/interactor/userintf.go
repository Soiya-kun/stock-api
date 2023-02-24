package interactor

import (
	"gitlab.com/soy-app/stock-api/domain/entity"
)

type UserCreate struct {
	Email    string
	UserType string
	Password string
	Name     string
}

type UserUpdate struct {
	UserID   string
	UserType string
	Name     string
}

type UserUpdatePassword struct {
	UserID      string
	NewPassword string
}

type IUserUseCase interface {
	Authenticate(token string) (string, error)
	Create(UserCreate) (entity.User, error)
	Delete(userID string) (entity.User, error)
	FindByID(userID string) (entity.User, error)
	FindByIDByAdmin(userID string) (entity.User, error)
	Login(email, password string) (entity.User, string, error)
	Search(query, userType string, skip int, limit int) ([]entity.User, int, error)
	SendResetPasswordMail(email string) error
	Update(UserUpdate) (entity.User, error)
	UpdatePassword(UserUpdatePassword) error
}
