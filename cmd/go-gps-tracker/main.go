package main

import (
	"flag"
	"fmt"

	"github.com/go-gps-tracker"
)

func main() {
	var db_host = flag.String("dbhost", "localhost", "Hostname of the InfluxDB server.")
	var db_port = flag.Int("dbport", 8086, "Port of the InfluxDB server")
	var db_user = flag.String("dbuser", "", "Username for the InfluxDB")
	var db_pass = flag.String("dbpass", "", "Password for the InfluxDB")
	var db_name = flag.String("dbname", "gps", "Name of the InfluxDB database")
	var port = flag.Int("port", 9000, "UDP port at which to listen for GPS traffic")
	flag.Parse()

	db := go_gps_tracker.DbConfig{
		//Address: "http://192.168.2.200:8186",
		Address:   fmt.Sprintf("https://%s:%d", *db_host, *db_port),
		Database:  *db_name,
		Precision: "s",
	}
	if *db_user != "" && *db_pass != "" {
		db.Username = *db_user
		db.Password = *db_pass
	}
	db.Connect()

	c := make(chan go_gps_tracker.Message, 100)
	quit := make(chan int)
	go go_gps_tracker.Listener(*port, c, quit)
	go db.Persist(c)
	<-quit
}
