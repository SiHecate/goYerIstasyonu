package ComPort

import (
	"fmt"
	"log"
	connect "yeristasyonu/ComPort/ConnectComPort"
	"yeristasyonu/Database/DatabaseConnection"
	databaseCreate "yeristasyonu/Database/DatabaseCreate"
)

func ComRead() {
	db, _ := DatabaseConnection.DatabaseConnect()
	databaseCreate.DbCreate()
	tableName := databaseCreate.GetCreatedTableName()

	for {
		buf := make([]byte, 128)               // Max data
		n, err := connect.SerialPort.Read(buf) // Seri portu kullanarak okuma
		if err != nil {
			log.Fatal(err)
		}
		data := string(buf[:n])
		query := fmt.Sprintf("INSERT INTO %s (data) VALUES ($1)", tableName)
		_, err = db.Exec(query, data)
		fmt.Println(data)
	}
}
