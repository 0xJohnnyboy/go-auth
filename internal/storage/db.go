package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	. "goauth/internal/models"
)

var dbPath = "goauth.db"

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&User{}, &RefreshToken{})
	return db, nil
}
