package service

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserService struct{}

// cardを指定してユーザ情報を取得
func (s UserService) GetUserByCard(db *gorm.DB, c echo.Context) ([]User, error) {
	var u []User
	card := c.Param("card")

	if err := db.Where("card = ?", card).First(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

// Post
// ユーザ追加
func (s UserService) PostUser(db *gorm.DB, c echo.Context) (User, error) {
	var user User
	c.Bind(&user)

	if err := db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
