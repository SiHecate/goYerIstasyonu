// main paketi
package main

import (
	ReadCom "yeristasyonu/ComPort/ComPortRead"

	_ "github.com/lib/pq"
)

func main() {
	ReadCom.ComRead()
}
