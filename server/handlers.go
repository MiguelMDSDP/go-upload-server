package server


import (
	"log"
	"net/http"
)


func UploadHandler(writer http.ResponseWriter, request *http.Request) {
	log.Print("**Upload Endpoint Hitted!")
}