package stats

import (
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
)

//CollectStat collects a stat
func CollectStat(measurement string, tags map[string]string, fields ...map[string]interface{}) {

	//Empty tags
	if tags == nil {
		tags = map[string]string{}
	}

	//Empty fields
	if fields == nil {
		fields = []map[string]interface{}{{"event": true}}
	}

	//Create point
	point := influxdb2.NewPoint(measurement, tags, fields[0], time.Now())

	//Write point
	writeAPI.WritePoint(point)
}
