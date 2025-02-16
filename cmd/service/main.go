package main

import (
	"log"
	"net/http"
	"merch-shop-service/internal/config"
)

func main() {
	db, err := config.OpenDB()
	if err != nil {
		log.Fatalf("Error to connect DB: %v", err)
	}
	defer db.Close()

	log.Println("The server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}