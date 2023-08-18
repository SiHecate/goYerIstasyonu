package ComPort

import (
	"fmt"
	"log"
)

func ComWrite() {
	if err := ConnectComPort(); err != nil {
		log.Fatal(err)
	}
	defer SerialPort.Close()

	var Send string
	fmt.Print(": ")
	fmt.Scanln(&Send)

	SerialPort.Write([]byte(Send))
}
