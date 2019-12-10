package main



import (
	"os"
	"fmt"
	"log"
	"github.com/MiguelMDSDP/upload-server/filesystem"
	"github.com/MiguelMDSDP/upload-server/server"
)



type UploadServer struct {
	// Native
	file_system *filesystem.FileSystem
	http_server *server.HTTPServer
	//Error Channels
	http_error_channel chan error
}



func NewUploadServer() *UploadServer {
	upload_server := new(UploadServer)

	// Native
	upload_server.file_system = filesystem.NewFileSystem()
	upload_server.http_server = server.NewHTTPServer()
	// Erros Channels
	upload_server.http_error_channel = make(chan error)

	return upload_server
}


func (upload_server *UploadServer) InitUploadServer() error {
	if err := upload_server.file_system.InitFileSystem(os.Getenv("HOME") + "/.ds"); err != nil {
		return err
	}
	if err := upload_server.http_server.InitHTTPServer(":8080"); err != nil {
		return err
	}
	return nil
}


func (upload_server *UploadServer) RunUploadServer() error {
	go upload_server.http_server.RunHTTPServer()
	
	log.Print("===== Upload Server Started =====")

	for {
		select {
			case err := <-upload_server.http_error_channel:
				if err != nil {
					return fmt.Errorf("%s: %s", "HTTP", err.Error())
				}
		}
	}
}


func (upload_server *UploadServer) StopUploadServer() {
	upload_server.file_system.StopFileSystem()
	upload_server.http_server.StopHTTPServer()
	close(upload_server.http_error_channel)
}