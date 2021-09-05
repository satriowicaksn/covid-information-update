package routes

import (
	"covid-information-update/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/data", controllers.GetData)
	e.GET("/data_log", controllers.GetDataByDayController)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST},
	}))
	return e
}
