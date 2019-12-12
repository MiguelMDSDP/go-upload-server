package main

import (
	"log"

	"github.com/MiguelMDSDP/upload-server/server"
)

func main() {
	httpServer := server.NewHTTPServer(":8080")
	if err := httpServer.Init(); err != nil {
		log.Fatal(err.Error())
	}
	if err := httpServer.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
