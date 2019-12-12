package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MiguelMDSDP/upload-server/filesystem"
)

// LsItem - ls route item response structure
type LsItem struct {
	FileName string `json:"fileName"`
	FileID   string `json:"fileID"`
}

// LsHandler is the function that treats the /ls route
func LsHandler(writer http.ResponseWriter, request *http.Request) {
	response := []LsItem{}

	for fileName, fileID := range filesystem.Indexes() {
		response = append(response, LsItem{fileName, fileID})
	}

	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

// UploadHandler is the function that treats the /upload route
func UploadHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("**Upload started.")

	err := request.ParseMultipartForm(32 << 20) // 32Mb
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	originalFile, handler, err := request.FormFile("file")
	defer originalFile.Close()
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	fileName := handler.Filename
	fileSize := handler.Size
	fileType := handler.Header.Get("Content-Type")

	log.Printf("-> File name: %v", fileName)
	log.Printf("-> File size: %v", fileSize)
	log.Printf("-> File type: %v", fileType)

	err = filesystem.CopyFile(originalFile, fileName)
	if err != nil {
		log.Println(err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("**Uploaded successfully!")
}
