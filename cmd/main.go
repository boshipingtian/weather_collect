package main

import (
	"time"
	"weather-colly/initial"
)

func main() {
	initial.Init()
	// Running forever
	timer := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-timer.C:
			timer.Reset(time.Second * 10)
		}
	}
}
