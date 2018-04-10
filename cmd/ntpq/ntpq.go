package main

import (
	"github.com/influxdata/telegraf/plugins/inputs/ntpq"
)

func main(){
	ntpq.Standalone()
}