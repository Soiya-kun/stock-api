package entity

import (
	"time"

	"gitlab.com/soy-app/stock-api/domain/entconst"
)

type User struct {
	UserID         string            `gorm:"primaryKey;size:30;not null"`
	Email          string            `gorm:"size:255;unique;not null;index"`
	HashedPassword string            `gorm:"size:255;not null"`
	UserType       entconst.UserType `gorm:"size:255;not null"`
	IsDeleted      bool              `gorm:"default:false"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (u User) IsSystemAdmin() bool {
	return u.UserType == entconst.SystemAdmin
}
