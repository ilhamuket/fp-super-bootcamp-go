package main

import (
	"final-project-golang-individu/config"
	"final-project-golang-individu/routes"
	"log"
)

func main() {
	// Inisialisasi database
	config.InitDB()

	// Inisialisasi router
	r := routes.SetupRouter()

	// Jalankan server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
