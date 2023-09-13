package main

import (
	"fmt"
	"polen/delivery"
	"time"
)

func main() {
	go delivery.NewServer().Run()

	// Jadwal pembaruan database setiap 24 jam
	updateInterval := 24 * time.Hour

	// Mulai penjadwalan pembaruan database
	ticker := time.NewTicker(updateInterval)

	// Loop tak terbatas untuk menjalankan pembaruan database
	for range ticker.C {
		// Panggil fungsi untuk melakukan pembaruan database di sini
		// usecase.UpdateDepositeDatabe()
		// Contoh: updateDatabase()
		fmt.Println("Melakukan pembaruan database...")
	}
}
