package ComPort

import (
	"fmt"
	"log"
	"yeristasyonu/Database/DatabaseConnection"
	databaseCreate "yeristasyonu/Database/DatabaseCreate"

	"github.com/tarm/serial"
)

func ConnectComPort() {

	c := &serial.Config{
		Name: "COM3",
		Baud: 9600,
	}

	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	db, _ := DatabaseConnection.DatabaseConnect()
	databaseCreate.DbCreate()
	tableName := databaseCreate.GetCreatedTableName()

	for {
		buf := make([]byte, 128) // Max data
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		data := string(buf[:n])
		query := fmt.Sprintf("INSERT INTO %s (data) VALUES ($1)", tableName)
		_, err = db.Exec(query, data)
		fmt.Println(data)
	}
}
