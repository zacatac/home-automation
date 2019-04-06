// +build darwin
package driver

import (
	"time"

	"github.com/sirupsen/logrus"
)

func NewLEDDriver() LEDDriver {
	return &driver{}
}

type driver struct{}

func (d *driver) DriveLED() error {
	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			<-ticker.C
			logrus.Infoln("blinking-led")
		}
	}()
	return nil
}
