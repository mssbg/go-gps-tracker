package go_gps_tracker

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/influxdata/influxdb/client/v2"
)

func error_check(err error) {
	if err != nil {
		fmt.Printf("error %s", err)
		log.Println("There was an error:", err)
	}
}

func error_fail(err error) {
	if err != nil {
		log.Fatalln("There was a fatal error:", err)
	}
}

type Coordinates struct {
	imei      string
	altitude  float64
	latitude  float64
	longitude float64
	time      string
}

type Message struct {
	msg    string
	size   int
	source string
}

type DbConfig struct {
	Address   string
	Username  string
	Password  string
	Database  string
	Precision string
	client    client.Client
}

func parse_ll(s string, n int, is_positive bool) (f float64) {
	a, err := strconv.ParseFloat(s[0:n+1], 64)
	error_check(err)
	d, err := strconv.ParseFloat(s[n+1:], 64)
	error_check(err)
	d /= 60.0
	res := a + d
	if is_positive {
		return res
	} else {
		return -res
	}
}

func parse_message(msg string) (c Coordinates) {
	fields := strings.Split(msg, ",")
	c.imei = strings.Split(fields[0], ":")[1]
	c.time = fields[2]
	//still not clear on what the format for this is. it doesn't look like meters.
	c.altitude, _ = strconv.ParseFloat(fields[5], 64)
	c.latitude = parse_ll(fields[7], 1, "N" == fields[8])
	c.longitude = parse_ll(fields[9], 2, "E" == fields[10])
	return c
}
