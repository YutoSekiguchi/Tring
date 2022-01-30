package controller

import (
	"gorm.io/gorm"
)

type Controller struct {
	Db *gorm.DB
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{
		Db: db,
	}
}