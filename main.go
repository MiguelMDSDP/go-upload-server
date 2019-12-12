package main

import "log"

func main() {
	uploadServer := NewUploadServer()
	if err := uploadServer.Init(); err != nil {
		log.Fatal(err.Error())
	}
	if err := uploadServer.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
