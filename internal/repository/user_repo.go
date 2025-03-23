package repository

import (
	"cuti/internal/model"

	"gorm.io/gorm"
)

func InsertUser(db *gorm.DB, user model.User) (model.User, error) {

	err := db.Create(&user).Error
	return user, err

}
