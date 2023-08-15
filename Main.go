// main paketi
package main

import (
	"yeristasyonu/ComPort"
	"yeristasyonu/Database/DatabaseConnection"

	_ "github.com/lib/pq"
)

func main() {
	DatabaseConnection.DatabaseConnect()
	ComPort.ConnectComPort()

}
