package stats

import (
	"os"

	"github.com/sirupsen/logrus"

	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
)

var (
	writeAPI api.WriteAPI
)

//Connect connects to influxdb using the `INFLUXDB_URL`, `INFLUXDB_TOKEN`, and `INFLUXDB_BUCKET` env variables
func Connect() {

	//Create influxdb client
	client := influxdb2.NewClient(os.Getenv("INFLUXDB_URL"), os.Getenv("INFLUXDB_TOKEN"))

	//Get write api
	writeAPI = client.WriteAPI("eutenly", os.Getenv("INFLUXDB_BUCKET"))

	//Log
	logrus.Info("InfluxDB is connected")
}
