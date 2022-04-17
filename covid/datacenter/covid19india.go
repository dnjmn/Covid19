package datacenter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type covid19india struct{}

type covidCases map[string]state

type state struct {
	Total struct {
		Confirmed int `json:"confirmed"`
	} `json:"total"`
}

func (c covid19india) fetchData() (CovidData, error) {

	data, err := requestForData()
	if err != nil {
		return nil, err
	}

	v := new(covidCases)
	err = json.Unmarshal(data, v)
	if err != nil {
		fmt.Println("error while fetching data from center: ", err)
		return nil, err
	}

	result := make(CovidData)
	for i, j := range *v {
		result[i] = j.Total.Confirmed
	}
	return result, nil
}

func requestForData() ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://data.covid19india.org/v4/min/data.min.json", nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return body, nil
}
