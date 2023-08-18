// main paketi
package main

import (
	Com "yeristasyonu/ComPort"

	_ "github.com/lib/pq"
)

func main() {
	Com.ComWrite()
}
