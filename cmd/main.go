package main

import (
	"log"

	"adapptor-backend/pkg/server"
)

func main() {
	server := server.New()

	if err := server.Start(8080); err != nil {
		log.Fatal(err)
	}
}
