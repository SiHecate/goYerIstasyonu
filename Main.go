// main paketi
package main

import (
	Database "yeristasyonu/Database/DatabaseList"

	_ "github.com/lib/pq"
)

func main() {
	Database.DbList()
}
