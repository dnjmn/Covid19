package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"projects.golang.dnjmn.com/covid-19/covid"
)

func main() {
	e := echo.New()
	e.GET("/", mainHandler)
	e.GET("/covid/case/update", covid.CovidUpdateHandler)
	e.POST("/covid/case/count/nearme", covid.GetCasesByLocation)
	e.Logger.Fatal(e.Start(":1323"))
}

func mainHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, dnjmn!")
}
