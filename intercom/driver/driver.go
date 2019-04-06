package driver

type LEDDriver interface {
	DriveLED() error
}
