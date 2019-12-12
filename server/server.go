package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HTTPServer - Type Definition
type HTTPServer struct {
	listeningAdress string
	router          *mux.Router
}

// NewHTTPServer creates HTTPServer objects
func NewHTTPServer() *HTTPServer {
	httpServer := new(HTTPServer)
	return httpServer
}

// Init initiates the HTTPServer
func (httpServer *HTTPServer) Init(address string) error {
	httpServer.router = mux.NewRouter()
	httpServer.router.HandleFunc("/upload", UploadHandler).Methods("POST")
	httpServer.listeningAdress = address
	return nil
}

// Run starts the HTTPServer execution
func (httpServer *HTTPServer) Run() error {
	err := http.ListenAndServe(httpServer.listeningAdress, httpServer.router)
	if err != nil {
		return err
	}
	return nil
}

// Stop stalls the HTTPServer execution
func (httpServer *HTTPServer) Stop() {
	log.Print("**Shutting down the http server.")
}
