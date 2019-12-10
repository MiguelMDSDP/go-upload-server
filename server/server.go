package server



import (
	"net/http"
	"github.com/gorilla/mux"
)



type HTTPServer struct {
	listening_adress string
	router *mux.Router
}



func NewHTTPServer() *HTTPServer {
	http_server := new(HTTPServer)
	return http_server
}


func (http_server *HTTPServer) InitHTTPServer(address string) error {
	http_server.router = mux.NewRouter()
	http_server.router.HandleFunc("/upload", UploadHandler).Methods("POST")
	return nil
}


func (http_server *HTTPServer) RunHTTPServer() error {
	err := http.ListenAndServe(http_server.listening_adress, http_server.router)
	if err != nil {
		return err
	}
	return nil
}