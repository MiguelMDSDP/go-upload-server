package main


import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
)


func uploadHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	request.ParseMultipartForm(10 << 20)

	file, handler, err := request.FormFile("myFile")
    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }
    defer file.Close()
    fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    tempFile, err := ioutil.TempFile("my-files", +"upload-test-*.pdf")
    if err != nil {
        fmt.Println(err)
    }
    defer tempFile.Close()

    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }
	
    tempFile.Write(fileBytes)
    
    fmt.Fprintf(writer, "Successfully Uploaded File\n")
}


func setupRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/upload", uploadHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}


func main() {
	fmt.Println("======ToT FileSystem Started======")
	setupRoutes()
}