package model

type Time struct {
	ID          int       `gorm:"primary_key;not null;autoIncrement:true"`
	UID         int       `gorm:"column:uid"`
	Time        int       `gorm:"column:time"`
	FinishAt    string    `gorm:"column:finish_at"`
	CreatedAt   string    `gorm:"column:created_at"`
}