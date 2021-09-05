package routes

import (
	"covid-information-update/controllers"
	"covid-information-update/database"
	"fmt"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func task() {
	fmt.Print("Doing task at")
	fmt.Print(time.Now())
	fmt.Println()
}

func New() *echo.Echo {
	e := echo.New()
	e.GET("/data", controllers.GetData)
	e.GET("/data_log", controllers.GetDataLogController)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST},
	}))

	q := make(chan bool)
	go Jobs(q)
	time.Sleep(2 * time.Second)
	q <- true
	close(q)

	return e
}

func Jobs(quit <-chan bool) {
	for {
		s := gocron.NewScheduler()
		s.Every(1).Day().At("01:26:00").Do(database.PostData)
		select {
		case <-quit:
			return
		case <-s.Start():
		}
	}
}
