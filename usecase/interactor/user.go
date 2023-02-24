package interactor

import (
	"errors"

	"gitlab.com/soy-app/stock-api/adapter/email"
	"gitlab.com/soy-app/stock-api/domain/constructor"
	"gitlab.com/soy-app/stock-api/domain/entity"
	"gitlab.com/soy-app/stock-api/usecase/port"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
)

type UserUseCase struct {
	email    port.Email
	ulid     port.ULID
	userAuth port.UserAuth
	userRepo port.UserRepository
}

func NewUserUseCase(
	email port.Email,
	ulid port.ULID,
	userAuth port.UserAuth,
	userRepo port.UserRepository,
) IUserUseCase {
	return &UserUseCase{
		email:    email,
		ulid:     ulid,
		userAuth: userAuth,
		userRepo: userRepo,
	}
}

func (u *UserUseCase) Login(email, password string) (entity.User, string, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return entity.User{}, "", err
	}

	err = u.userAuth.CheckPassword(user, password)
	if err != nil {
		return entity.User{}, "", err
	}

	token, err := u.userAuth.IssueUserToken(user)
	if err != nil {
		return entity.User{}, "", err
	}

	return user, token, nil
}

func (u *UserUseCase) Authenticate(token string) (string, error) {
	return u.userAuth.Authenticate(token)
}

func (u *UserUseCase) FindByID(userID string) (entity.User, error) {
	return u.userRepo.FindByID(userID)
}

func (u *UserUseCase) FindByIDByAdmin(userID string) (entity.User, error) {
	return u.userRepo.FindByID(userID)
}

func (u *UserUseCase) Create(userCreate UserCreate) (entity.User, error) {
	_, err := u.userRepo.FindByEmail(userCreate.Email)
	if err == ErrUserNotFound {
		// errがErrUserNotFoundの場合はcreateできるのでreturnしない
	} else if err != nil {
		// errがErrUserNotFound以外の場合はそのerrを返す
		return entity.User{}, err
	} else {
		return entity.User{}, ErrUserAlreadyExists
	}

	hp, err := u.userAuth.HashPassword(userCreate.Password)
	if err != nil {
		return entity.User{}, err
	}

	userID := u.ulid.New()
	user, err := constructor.NewUserCreate(
		userCreate.Email,
		hp,
		userCreate.Password,
		userID,
		userCreate.UserType,
	)
	if err != nil {
		return entity.User{}, err
	}

	err = u.userRepo.Create(user)
	if err != nil {
		return entity.User{}, err
	}

	return u.userRepo.FindByID(user.UserID)
}

func (u *UserUseCase) Search(query string, userType string, skip int, limit int) ([]entity.User, int, error) {
	return u.userRepo.Search(query, userType, skip, limit)
}

func (u *UserUseCase) Update(user UserUpdate) (entity.User, error) {
	// TODO: 更新処理
	return u.userRepo.FindByID(user.UserID)
}

func (u *UserUseCase) Delete(userID string) (entity.User, error) {
	res, err := u.userRepo.FindByID(userID)
	if err != nil {
		return entity.User{}, err
	}

	err = u.userRepo.Delete(userID)
	if err != nil {
		return entity.User{}, err
	}

	return res, nil
}

func (u *UserUseCase) SendResetPasswordMail(emailAddress string) error {
	user, err := u.userRepo.FindByEmail(emailAddress)
	if err != nil {
		return err
	}

	token, err := u.userAuth.IssueUserToken(user)
	if err != nil {
		return err
	}

	subject, body := email.ContentToResetPassword(token)
	err = u.email.Send([]string{emailAddress}, subject, body, "")
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUseCase) UpdatePassword(userUpdatePassword UserUpdatePassword) error {
	user, err := u.userRepo.FindByID(userUpdatePassword.UserID)
	if err != nil {
		return err
	}

	hashedPassword, err := u.userAuth.HashPassword(userUpdatePassword.NewPassword)
	if err != nil {
		return err
	}

	user.HashedPassword = hashedPassword
	return u.userRepo.UpdatePassword(user)
}
