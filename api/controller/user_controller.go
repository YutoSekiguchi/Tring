package controller

import (
	"fmt"
	"net/http"

	"github.com/YutoSekiguchi/tring/service"
	"github.com/labstack/echo/v4"
)

// HandleGetUserByID GET /users/:card idが一致するユーザを取得
func (ctrl Controller) HandleGetUserByCard(c echo.Context) error {
	var s service.UserService
	p, err := s.GetUserByCard(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(200, nil)
	} else {
		return c.JSON(200, p)
	}
}


// HandlePostUser POST /users ユーザの追加
func (ctrl Controller) HandlePostUser(c echo.Context) error {
	var s service.UserService
	p, err := s.PostUser(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}