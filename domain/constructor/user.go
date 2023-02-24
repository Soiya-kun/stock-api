package constructor

import (
	"gitlab.com/soy-app/stock-api/domain/entconst"
	"gitlab.com/soy-app/stock-api/domain/entity"
	"gitlab.com/soy-app/stock-api/domain/validation"
)

func NewUserCreate(
	email string,
	hashedPassword string,
	password string,
	userID string,
	userType string,
) (entity.User, error) {
	err := validation.ValidateEmail(email)
	if err != nil {
		return entity.User{}, err
	}

	err = validation.ValidatePassword(password)
	if err != nil {
		return entity.User{}, err
	}

	userTypeConverted, err := entconst.UserTypeFromString(userType)
	if err != nil {
		return entity.User{}, err
	}

	return entity.User{
		UserID:         userID,
		Email:          email,
		UserType:       *userTypeConverted,
		HashedPassword: hashedPassword,
	}, nil
}

func NewUserUpdate(
	email string,
	userID string,
	userType string,
) (entity.User, error) {
	err := validation.ValidateEmail(email)
	if err != nil {
		return entity.User{}, err
	}

	userTypeValidated, err := entconst.UserTypeFromString(userType)
	if err != nil {
		return entity.User{}, err
	}

	return entity.User{
		UserID:         userID,
		Email:          email,
		UserType:       *userTypeValidated,
		HashedPassword: "",
	}, nil
}
