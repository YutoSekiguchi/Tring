package model

import (
	"time"
)

type User struct {
	ID          int       `gorm:"primary_key;not null;autoIncrement:true"`
	Name        string    `gorm:"type:text"`
	Image       string    `gorm:"type:text"`
	Card        string    `gorm:"type:text;not null"`
	CreatedAt   time.Time `sql:"DEFAULT:current_timestamp;column:created_at"`
}