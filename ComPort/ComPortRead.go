package ComPort

import (
	"fmt"
	"log"
	"yeristasyonu/Database/DatabaseConnection"
	databaseCreate "yeristasyonu/Database/DatabaseCreate"
)

func ComRead() {
	db, _ := DatabaseConnection.DatabaseConnect()
	databaseCreate.DbCreate()
	tableName := databaseCreate.GetCreatedTableName()

	if err := ConnectComPort(); err != nil {
		log.Fatal(err) // Olası bir hata için log yazdırma
	}
	defer SerialPort.Close()

	for {
		buf := make([]byte, 256) // Max data
		n, err := SerialPort.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		data := string(buf[:n])
		query := fmt.Sprintf("INSERT INTO %s (data) VALUES ($1)", tableName)
		_, err = db.Exec(query, data)
		fmt.Println(data)
	}
}
