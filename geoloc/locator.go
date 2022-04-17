package geoloc

import (
	"fmt"
	"os"
	"strings"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/opencage"
)

const (
	addr         = "Melbourne VIC"
	radius       = 50
	zoom         = 18
	addrFR       = "Champs de Mars Paris"
	latFR, lngFR = 12.9882, 77.6121
)

func Geo(l1, l2 float64) (string, *geo.Address) {
	code := os.Getenv("OPENCAGE_API_KEY")
	geocoder := opencage.Geocoder(code)
	address, err := geocoder.ReverseGeocode(l1, l2)
	if err != nil {
		fmt.Println("error geocoding", l1, l2)
	}
	fmt.Println(address.Country, address.State, address.StateCode)
	if code, ok := stateToCode[strings.ToLower(address.State)]; ok {
		return code, address
	}
	return "", nil
}
