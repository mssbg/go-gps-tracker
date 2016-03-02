# go-gps-tracker

## Overview

A simple server for collecting TK103B GPS tracker data

Currently testedonly with TK103B, but possibly it will work with other models from this series.
It will listen for UDP traffic and store each message in InfluxDB database.
If you have any other similar model, send me captured traffic from it and I'll add support for it.

## Usage

```
$ ./go-gps-tracker --help
Usage of ./go-gps-tracker:
  -dbhost string
    	Hostname of the InfluxDB server. (default "localhost")
  -dbname string
    	Name of the InfluxDB database (default "gps")
  -dbpass string
    	Password for the InfluxDB
  -dbport int
    	Port of the InfluxDB server (default 8086)
  -dbuser string
    	Username for the InfluxDB
  -port int
    	UDP port at which to listen for GPS traffic (default 9000)
```

The program doesn't have functionality to detach from the terminal. Use `screen`, `tmux` or `systemd` to run it as a service.

## Building from sources
### Pre-requisites
* Go language runtime 1.5 or later
* Glide dependency management for Go

### Build workflow
```
$ git clone https://github.com/mssbg/go-gps-tracker.git
$ cd go-gps-tracker
$ glide install
```

