package server

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

// UploadHandler is the function that treats the /upload route
func UploadHandler(writer http.ResponseWriter, request *http.Request) {
	log.Print("**Starting upload")
	// multipartReader, err := request.MultipartReader()
	// if err != nil {
	// 	log.Print(err.Error())
	// 	writer.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	file, _, err := request.FormFile("file")
	defer file.Close()
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	path := filepath.Join(".data/", uuid.New().String())
	f, err := os.Create(path)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	f, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	io.Copy(f, file)
	// buffer := make([]byte, 100000)
	// var cBytes int
	// for {
	// 	part, err := multipartReader.NextPart()
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			log.Println(err.Error())
	// 			file.Close()
	// 			break
	// 		}
	// 		log.Println(err.Error())
	// 		err = os.Remove(path)
	// 		if err != nil {
	// 			log.Println(err.Error())
	// 		}
	// 		writer.WriteHeader(http.StatusBadRequest)
	// 		return
	// 	}
	// 	cBytes, err = part.Read(buffer)
	// 	if err != nil {
	// 		log.Println(err.Error)
	// 		writer.WriteHeader(http.StatusBadRequest)
	// 		return
	// 	}
	// 	_, err = file.Write(buffer[0:cBytes])
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		err = os.Remove(path)
	// 		if err != nil {
	// 			log.Println(err.Error())
	// 		}
	// 		writer.WriteHeader(http.StatusInternalServerError)
	// 		return
	// 	}
	// 	file.Sync()
	// }
}
