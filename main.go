package main

import (
	"go-crud/config"
	"go-crud/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"Hello": "world",
		})
	})

	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

	//rout api
	route := e.Group("/user")
	route.GET("/", controller.GetUser)
	route.GET("/:id", controller.GetById)
	route.POST("/", controller.AddUser)
	route.PUT("/:id", controller.UpdateUser)
	route.DELETE("/:Id", controller.DeleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}
