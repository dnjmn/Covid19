package covid

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"projects.golang.dnjmn.com/covid-19/geoloc"
)

// Location type has latitude and longitude
type Location struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

// CovidUpdateHandler fetches and updates latest total positive cases in india
func CovidUpdateHandler(c echo.Context) error {
	count, err := UpdateData()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error occured")
	}
	scount := strconv.Itoa(count)
	return c.JSON(http.StatusOK, "updated cases: "+scount)
}

// GetCasesByLocation accepts the location and returns the total count of positive cases in that state
func GetCasesByLocation(c echo.Context) error {
	v := new(Location)
	if err := c.Bind(v); err != nil {
		return err
	}

	code, address := geoloc.Geo(v.Latitude, v.Longitude)
	if code == "" {
		return c.JSON(http.StatusNotFound, "covid data not found for given location")
	}
	covidData := GetByStateCode(code)
	return c.JSON(http.StatusOK, "total positive covid count cases in "+address.State+" is: "+strconv.Itoa(covidData.PositiveCount))
}
