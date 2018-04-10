package main

import (
	"github.com/influxdata/telegraf/plugins/inputs/ntpq"
	"fmt"
)

func main(){
	fmt.Println("start")
	ntpq.Standalone()
}