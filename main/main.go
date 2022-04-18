package main

import (
	"Covid19/covid"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", mainHandler)
	e.GET("/covid/case/update", covid.CovidUpdateHandler)
	e.POST("/covid/case/count/nearme", covid.GetCasesByLocation)
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("port is empty")
		return
	}
	e.Logger.Fatal(e.Start(":" + port))
}

func mainHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, dnjmn!")
}
