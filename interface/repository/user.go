package repository

import (
	"gorm.io/gorm"

	"gitlab.com/soy-app/stock-api/domain/entity"
	"gitlab.com/soy-app/stock-api/usecase/interactor"
	"gitlab.com/soy-app/stock-api/usecase/port"
)

type UserRepository struct {
	db   *gorm.DB
	ulid port.ULID
}

func NewUserRepository(
	db *gorm.DB,
	ulid port.ULID,
) port.UserRepository {
	return &UserRepository{db: db, ulid: ulid}
}

func (r *UserRepository) Create(user entity.User) error {
	m := &entity.User{
		UserID:         user.UserID,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		IsDeleted:      false,
	}
	if err := r.db.Create(m).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Delete(userID string) error {
	return r.db.Model(&entity.User{}).
		Where("user_id = ?", userID).
		Updates(
			map[string]interface{}{
				"email":           userID + "_deleted",
				"hashed_password": "",
				"is_deleted":      true,
			},
		).Error
}

func (r *UserRepository) FindByEmail(email string) (entity.User, error) {
	ret := entity.User{}
	err := r.db.Where("email = ?", email).First(&ret).Error
	if err == gorm.ErrRecordNotFound {
		return entity.User{}, interactor.ErrUserNotFound
	}
	if err != nil {
		return entity.User{}, err
	}
	return ret, nil
}

func (r *UserRepository) FindByID(userID string) (entity.User, error) {
	res := entity.User{}
	err := r.db.
		Preload("UserDetail").
		Preload("UserDetail").
		Where("is_deleted = false").
		Where("user_id = ?", userID).
		First(&res).
		Error
	if err == gorm.ErrRecordNotFound {
		return entity.User{}, interactor.ErrUserNotFound
	}
	if err != nil {
		return entity.User{}, err
	}
	return res, nil
}

func (r *UserRepository) Search(query string, userType string, skip int, limit int) ([]entity.User, int, error) {
	var res []entity.User
	var total int64
	sqlQuery := r.db.Model(&entity.User{}).
		Preload("UserDetail").
		Joins("LEFT OUTER JOIN user_details ON user_details.user_id = users.user_id").
		Where("email LIKE ?", "%"+query+"%").
		Where("is_deleted = false")

	if userType != "" {
		sqlQuery = sqlQuery.Where("user_type = ?", userType)
	}

	err := sqlQuery.
		Group("users.user_id").
		Count(&total).
		Offset(skip).
		Limit(limit).
		Find(&res).
		Error
	if err != nil {
		return nil, 0, err
	}

	return res, int(total), nil
}

func (r *UserRepository) Update(user entity.User) error {
	return r.db.Model(&entity.User{}).
		Where("user_id = ?", user.UserID).
		Updates(
			map[string]interface{}{
				"email":     user.Email,
				"user_type": user.UserType,
			},
		).Error
}

func (r *UserRepository) UpdatePassword(user entity.User) error {
	return r.db.Model(&entity.User{}).
		Where("user_id = ?", user.UserID).
		Updates(
			map[string]interface{}{
				"hashed_password": user.HashedPassword,
			},
		).Error
}
