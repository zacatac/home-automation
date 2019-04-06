// +build linux
package driver

import (
	"time"

	"github.com/sirupsen/logrus"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func NewLEDDriver() LEDDriver {
	return &driver{}
}

type driver struct{}

func (d *driver) DriveLED() error {
	r := raspi.NewAdaptor()
	led := gpio.NewLedDriver(r, "7")

	work := func() {
		gobot.Every(1*time.Second, func() {
			logrus.Infoln("blinking-led")
			if err := led.Toggle(); err != nil {
				logrus.WithField("error", err).Errorln("blinking-led-failed")
			}

		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{led},
		work,
	)

	return robot.Start()
}
