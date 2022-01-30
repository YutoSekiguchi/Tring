package service

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TimeService struct{}

type TimeUser struct {
	ID          int       `gorm:"primary_key;not null;autoIncrement:true"`
	Name        string
	Card        string
	Image       string
	UID         int       `gorm:"column:uid"`
	Time        int       `gorm:"column:time"`
	FinishAt    string    `gorm:"column:finish_at"`
	CreatedAt   string    `gorm:"column:created_at"`
}

// 休憩時間とユーザ情報一覧を取得
func (s TimeService) GetTimeList(db *gorm.DB, c echo.Context) ([]TimeUser, error) {
	var tu []TimeUser
	
	if err := db.Debug().Raw("SELECT times.id, users.name, users.card, users.image, times.uid, times.time, times.finish_at, times.created_at FROM times JOIN users ON times.uid = users.id ORDER BY times.finish_at ASC").Scan(&tu).Error; err != nil {
		return nil, err
	}
	return tu, nil
}

// uidを指定して休憩時間の取得
func (s TimeService) GetTimeByUID(db *gorm.DB, c echo.Context) ([]Time, error) {
	var t []Time
	uid := c.Param("uid")

	if err := db.Where("uid = ?", uid).First(&t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

// idを指定して休憩時間の削除
func (s TimeService) DeleteTime(db *gorm.DB, c echo.Context) ([]Time, error) {
	var t []Time
	id := c.Param("id")

	if err := db.Where("id = ?", id).Delete(&t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

// 休憩時間の追加
func (s TimeService) PostTime(db *gorm.DB, c echo.Context) (Time, error) {
	var t Time
	c.Bind(&t)

	if err := db.Table("times").Create(&t).Error; err != nil {
		return t, err
	}
	return t, nil
}

// uidを指定してtimeを削除
func (s TimeService) DeleteTimeByUID(db *gorm.DB, c echo.Context) ([]Time, error) {
	var t []Time
	uid := c.Param("uid")
	
	if err := db.Where("uid = ?", uid).Delete(&t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

