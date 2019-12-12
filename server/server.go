package server

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// HTTPServer - Type Definition
type HTTPServer struct {
	listeningAdress string
	router          *mux.Router
}

// NewHTTPServer creates HTTPServer objects
func NewHTTPServer(address string) *HTTPServer {
	httpServer := new(HTTPServer)

	httpServer.router = mux.NewRouter()
	httpServer.listeningAdress = address

	return httpServer
}

// Init initiates the HTTPServer
func (httpServer *HTTPServer) Init() error {
	httpServer.router.HandleFunc("/upload", UploadHandler).Methods("POST")
	httpServer.router.HandleFunc("/ls", LsHandler).Methods("GET")

	fileServer := http.FileServer(http.Dir("data/"))
	httpServer.router.Handle("/download/{id}", http.StripPrefix("/download", fileServer))

	return nil
}

// Run starts the HTTPServer execution
func (httpServer *HTTPServer) Run() error {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Println("===== ToTFS Running =====")
	err := http.ListenAndServe(httpServer.listeningAdress, handlers.CORS(headersOk, originsOk, methodsOk)(httpServer.router))
	if err != nil {
		return err
	}

	return nil
}

// Stop stalls the HTTPServer execution
func (httpServer *HTTPServer) Stop() {
	log.Print("**Shutting down the http server.")
}
