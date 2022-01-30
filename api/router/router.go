package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"

	"github.com/YutoSekiguchi/tring/controller"
)

func InitRouter(db *gorm.DB) {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", "http://localhost:7120"},
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	ctrl := controller.NewController(db)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go!")
	})

	// User
	user := e.Group("/users")
	{
		user.GET("/card/:card", ctrl.HandleGetUserByCard)
		user.POST("", ctrl.HandlePostUser)
	}

	// Time
	time := e.Group("/times")
	{
		time.GET("/list", ctrl.HandleGetTimeList)
		time.GET("/time/:uid", ctrl.HandleGetTimeByUID)
		time.DELETE("/delete/:id", ctrl.HandleDeleteTime)
		time.POST("", ctrl.HandlePostTime)
		time.DELETE("/usertime/delete/:uid", ctrl.HandleDeleteTimeByUID)
	}

	e.Logger.Fatal(e.Start(":9000"))
}