package repository

import (
	"cuti/internal/model"
	"fmt"

	"gorm.io/gorm"
)

func InsertUser(db *gorm.DB, user model.User) model.User {
	if err := db.Create(&user).Error; err != nil {
		fmt.Println("gagal membuat user", err)

	}
	defer db.Close()
	return user

}
