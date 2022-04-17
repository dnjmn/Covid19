package datacenter

var defaultDataCenter DataCenter

func init() {
	defaultDataCenter = new(covid19india)
}

func GetData() (CovidData, error) {
	return defaultDataCenter.fetchData()
}

type CovidData map[string]int

type DataCenter interface {
	fetchData() (CovidData, error)
}
