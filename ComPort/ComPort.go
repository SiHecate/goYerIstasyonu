package ComPort

import (
	"fmt"
	"log"
	"yeristasyonu/Database/DatabaseConnection"

	"github.com/tarm/serial"
)

func ConnectComPort() {

	// Seri port ayarları
	c := &serial.Config{
		Name: "COM3", // Kullandığınız portun adını burada belirtin
		Baud: 9600,   // İletişim hızını burada belirtin
	}

	// Seri port açma
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	db, err := DatabaseConnection.DatabaseConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Sonsuz döngüde seri porttan veri okuma
	for {
		buf := make([]byte, 128) // Okunacak maksimum veri boyutunu burada belirtin
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		data := string(buf[:n])
		_, err = db.Exec("INSERT INTO data_table (data) VALUES ($1)", data)
		if err != nil {
			log.Println("Veri kaydetme hatası:", err)
		} else {
			fmt.Println("Veri başarıyla kaydedildi:", data)
		}
	}

}
