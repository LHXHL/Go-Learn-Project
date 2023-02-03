package model

import (
	"LHXHL/go-mux-gorm-crud/internal/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Age      string `json:"age"`
	Password string `json:"password"`
}

func InitDB() {
	config.Connect()
	db = config.GetDB()
	_ = db.AutoMigrate(&User{})
}

func GetAllUsers() ([]User, error) {
	var user []User
	err := db.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserById(id int) (User, error) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(user User) (User, error) {
	err := db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUser(id int, user User) (User, error) {
	err := db.Model(&user).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func DeleteUser(id int) error {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
