package models

import (
	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model
	UserID    uint   `gorm:"index;constraint:OnDelete:CASCADE"`
	TokenHash string `gorm:"unique"`
	ExpiresAt int64
}
