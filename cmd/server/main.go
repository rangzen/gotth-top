package main

import "github.com/rangzen/gotha-top/handlers"

func main() {
	server := handlers.NewServer()
	server.Run()
}
