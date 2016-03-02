package go_gps_tracker

import (
	"log"
	"strings"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

func (db DbConfig) process_message(msg string, addr string) {
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "moto",
		Precision: "s",
	})
	if strings.HasPrefix(msg, "imei") {
		coordinates := parse_message(msg)
		tags := map[string]string{}
		fields := map[string]interface{}{
			"imei":        coordinates.imei,
			"altitude":    coordinates.altitude,
			"latitude":    coordinates.latitude,
			"longitude":   coordinates.longitude,
			"local_time":  coordinates.time,
			"raw":         msg,
			"source_addr": addr,
		}
		pt, err := client.NewPoint("coordinates", tags, fields, time.Now())
		error_check(err)
		bp.AddPoint(pt)
		err = db.client.Write(bp)
		error_fail(err)
	}
}

func (db *DbConfig) Connect() {
	var err error
	db.client, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     db.Address,
		Username: db.Username,
		Password: db.Password,
	})
	error_fail(err)
	log.Println("Conected to DB.")
}

func (db DbConfig) Persist(c chan Message) {
	for {
		m, ok := <-c
		if ok {
			log.Printf("Storing %s from %s", m.msg, m.source)
			db.process_message(m.msg, m.source)
		} else {
			break
		}
	}
}
