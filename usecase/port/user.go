package port

import (
	"gitlab.com/soy-app/stock-api/domain/entity"
)

type UserRepository interface {
	Create(entity.User) error
	Delete(userID string) error
	FindByEmail(email string) (entity.User, error)
	FindByID(userID string) (entity.User, error)
	Search(Query string, userType string, Skip int, Limit int) ([]entity.User, int, error)
	Update(entity.User) error
	UpdatePassword(entity.User) error
}
type UserAuth interface {
	Authenticate(token string) (string, error)
	CheckPassword(user entity.User, password string) error
	HashPassword(password string) (string, error)
	IssueUserToken(entity.User) (string, error)
	GenerateInitialPassword(length int) (string, error)
}
