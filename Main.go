// main paketi
package main

import (
	comport "yeristasyonu/ComPort"

	_ "github.com/lib/pq"
)

func main() {
	comport.ConnectComPort()
}
