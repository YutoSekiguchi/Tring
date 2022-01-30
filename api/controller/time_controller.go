package controller

import (
	"fmt"
	"net/http"

	"github.com/YutoSekiguchi/tring/service"
	"github.com/labstack/echo/v4"
)

// HandleGetTimeList GET /times/list
func (ctrl Controller) HandleGetTimeList(c echo.Context) error {
	var s service.TimeService
	p, err := s.GetTimeList(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandleGetTimeByUID GET /times/time/:uid
func (ctrl Controller) HandleGetTimeByUID(c echo.Context) error {
	var s service.TimeService
	p, err := s.GetTimeByUID(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(200, nil)
	} else { 
		return c.JSON(200, p)
	}
} 

// HandleDeleteTime DELETE /times/delete/:id
func (ctrl Controller) HandleDeleteTime(c echo.Context) error {
	var s service.TimeService
	p, err := s.DeleteTime(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandlePostTime POST /times
func (ctrl Controller) HandlePostTime(c echo.Context) error {
	var s service.TimeService
	p, err := s.PostTime(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandleDeleteTimeByUID DELETE /times/card/delete/:uid
func (ctrl Controller) HandleDeleteTimeByUID(c echo.Context) error {
	var s service.TimeService
	p, err := s.DeleteTimeByUID(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}