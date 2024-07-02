package main

import (
	"log"

	"github.com/rangzen/gotth-top/handlers"
)

func main() {
	server := handlers.NewServer()
	log.Fatal(server.Run())
}
