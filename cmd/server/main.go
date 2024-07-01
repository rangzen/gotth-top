package main

import "github.com/rangzen/gotth-top/handlers"

func main() {
	server := handlers.NewServer()
	server.Run()
}
