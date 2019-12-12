package main

import (
	"fmt"
	"log"

	"github.com/MiguelMDSDP/upload-server/filesystem"
	"github.com/MiguelMDSDP/upload-server/server"
)

// UploadServer - Upload server data structure
type UploadServer struct {
	// Native
	fileSystem *filesystem.FileSystem
	httpServer *server.HTTPServer
	//Error Channels
	httpErrChan chan error
}

// NewUploadServer creates new UploadServer objects
func NewUploadServer() *UploadServer {
	uploadServer := new(UploadServer)

	// Native
	uploadServer.fileSystem = filesystem.NewFileSystem()
	uploadServer.httpServer = server.NewHTTPServer()
	// Erros Channels
	uploadServer.httpErrChan = make(chan error)

	return uploadServer
}

// Init initiates the UploadServer object
func (uploadServer *UploadServer) Init() error {
	if err := uploadServer.fileSystem.Init(".data"); err != nil {
		return err
	}
	if err := uploadServer.httpServer.Init(":8080"); err != nil {
		return err
	}
	return nil
}

// Run starts the UploadServer execution
func (uploadServer *UploadServer) Run() error {
	go uploadServer.httpServer.Run()

	log.Print("===== Upload Server Started =====")

	for {
		select {
		case err := <-uploadServer.httpErrChan:
			if err != nil {
				return fmt.Errorf("%s: %s", "HTTP", err.Error())
			}
		}
	}
}

// Stop stalls the UploadServer execution
func (uploadServer *UploadServer) Stop() {
	uploadServer.fileSystem.Stop()
	uploadServer.httpServer.Stop()
	close(uploadServer.httpErrChan)
}
