package covid

import (
	"Covid19/covid/datacenter"
	"Covid19/db"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CovidDataCollection = "CovidDataCollection"
)

type CovidData struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	State         string             `bson:"state,omitempty"`
	PositiveCount int                `bson:"positiveCount,omitempty"`
}

func GetByStateCode(code string) (v CovidData) {
	client, ctx, cancel := db.GetClient()
	defer cancel()
	db := client.Database(db.Database)
	collection := db.Collection(CovidDataCollection)
	v.State = code
	result := collection.FindOne(ctx, &v)
	result.Decode(&v)

	return v
}

func UpdateData() (int, error) {
	client, ctx, cancel := db.GetClient()
	defer cancel()

	db := client.Database(db.Database)
	collection := db.Collection(CovidDataCollection)
	err := collection.Drop(ctx)
	if err != nil {
		log.Fatal("error while dropping collection", err)
	}

	result, err := collection.InsertMany(ctx, GetCovidData())
	if err != nil {
		log.Println(err)
		return 0, errors.New("error while inserting in db")
	}
	return len(result.InsertedIDs), nil
}

func GetCovidData() (data []interface{}) {
	vals, err := datacenter.GetData()
	if err != nil {
		log.Fatal(err)
	}
	for i, j := range vals {
		data = append(data, CovidData{
			State:         i,
			PositiveCount: j,
		})
	}
	return
}
