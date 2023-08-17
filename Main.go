// main paketi
package main

import (
	databaseCreate "yeristasyonu/Database/DatabaseCreate"

	_ "github.com/lib/pq"
)

func main() {
	databaseCreate.DbCreate()
}
