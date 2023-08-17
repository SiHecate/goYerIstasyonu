// main paketi
package main

import (
	ReadCom "yeristasyonu/ComPort"

	_ "github.com/lib/pq"
)

func main() {
	ReadCom.ComRead()
}
