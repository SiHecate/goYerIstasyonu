package databaseCreate

import (
	"fmt"
	"strings"
	"time"
	"yeristasyonu/Database/DatabaseConnection"

	_ "github.com/lib/pq"
)

func database_name_time() time.Time {
	return time.Now()
}

func database_name_name() string {
	var dbName string
	fmt.Print("Enter database name: ")
	fmt.Scan(&dbName)
	return dbName
}

func database_name_code() string {
	db, _ := DatabaseConnection.DatabaseConnect()

	db.Exec("CREATE SEQUENCE database_name_code START 1 INCREMENT 1;")

	var newCode string
	db.QueryRow("SELECT LPAD(NEXTVAL('database_name_code')::text, 3, '0')").Scan(&newCode)

	return newCode
}

func DbCreate() {
	db, err := DatabaseConnection.DatabaseConnect()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	time := database_name_time()
	name := database_name_name()
	code := database_name_code()

	nameWithoutSpaces := strings.ReplaceAll(name, " ", "_")

	tableName := fmt.Sprintf("%s_%s_%s", code, nameWithoutSpaces, time.Format("20060102150405"))

	query := fmt.Sprintf("CREATE TABLE \"%s\" (id SERIAL PRIMARY KEY, name TEXT)", tableName)
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	fmt.Printf("Table created with name: %s\n", tableName)
}
