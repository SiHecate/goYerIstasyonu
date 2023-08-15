// DatabaseConnection paketi
package DatabaseConnection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func DatabaseConnect() (*sql.DB, error) {
	databaseBilgileri := "user=postgres password=393406 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", databaseBilgileri)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Database'e giriş başarılı.")
	return db, nil
}
