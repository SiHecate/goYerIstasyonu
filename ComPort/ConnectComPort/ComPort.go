package ComPort

import (
	"log"

	"github.com/tarm/serial"
)

var SerialPort *serial.Port

func ConnectComPort() {

	c := &serial.Config{
		Name: "COM3",
		Baud: 9600,
	}

	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	SerialPort = s
}
