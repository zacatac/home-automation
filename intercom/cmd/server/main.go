package main

import (
	"github.com/sirupsen/logrus"
	"github.com/zacatac/home-automation/intercom/driver"
)

func main() {
	ledDriver := driver.NewLEDDriver()
	if err := ledDriver.DriveLED(); err != nil {
		logrus.Fatalln(err)
	}
}
