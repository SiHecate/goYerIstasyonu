package ComPort

import (
	"github.com/tarm/serial"
)

var SerialPort *serial.Port

func ConnectComPort() error {
	c := &serial.Config{
		Name: "COM3",
		Baud: 9600,
	}

	s, err := serial.OpenPort(c)
	if err != nil {
		return err
	}
	SerialPort = s
	return nil
}
