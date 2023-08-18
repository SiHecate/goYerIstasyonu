package Database

import (
	"fmt"
	DatabaseConnection "yeristasyonu/Database/DatabaseConnection"

	_ "github.com/lib/pq"
)

func DbList() {
	db, err := DatabaseConnection.DatabaseConnect()
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE';")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Tables:")
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(tableName)
	}
}
