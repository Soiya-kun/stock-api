package schema

import "gitlab.com/soy-app/stock-api/domain/entity"

type UserRes struct {
	UserId   string `json:"user_id"`
	Email    string `json:"email"`
	UserType string `json:"user_type"`
}

type UsersRes struct {
	List  []UserRes `json:"list"`
	Total int       `json:"total"`
}

type CreateUserReq struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

type CreateUserReqByAdmin struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

type UpdateUserReq struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	UserID   string `param:"user-id"`
	UserType string `json:"user_type"`
}

type UserSearchQueryReq struct {
	Query    string `query:"q"`
	UserType string `query:"user-type"`
	Skip     int    `query:"skip"`
	Limit    int    `query:"limit"`
}

func UserResFromEntity(user entity.User) UserRes {
	return UserRes{
		UserId:   user.UserID,
		Email:    user.Email,
		UserType: string(user.UserType),
	}
}

func UsersResFromSearchResult(list []entity.User, total int) UsersRes {
	users := make([]UserRes, len(list))
	for i, user := range list {
		users[i] = UserResFromEntity(user)
	}
	return UsersRes{
		List:  users,
		Total: total,
	}
}
