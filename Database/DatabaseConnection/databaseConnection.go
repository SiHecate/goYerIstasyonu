// DatabaseConnection paketi
package DatabaseConnection

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DatabaseConnect() (*sql.DB, error) {
	databaseInfo := "user=postgres password=393406 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", databaseInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
