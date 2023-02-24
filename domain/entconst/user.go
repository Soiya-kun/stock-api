package entconst

import "errors"

var (
	ErrInvalidUserType = errors.New("invalid user type")
)

type UserType string

const (
	SystemAdmin UserType = "SystemAdmin" // システム管理者
	User        UserType = "User"        // 一般ユーザー
)

func (u UserType) String() string {
	return string(u)
}

func UserTypeFromString(userType string) (*UserType, error) {
	ret := UserType(userType)
	switch ret {
	case SystemAdmin, User:
		return &ret, nil
	default:
		return nil, ErrInvalidUserType
	}
}
