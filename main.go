package main



import "log"


func main() {
	upload_server := NewUploadServer()
	if err := upload_server.InitUploadServer(); err != nil {
		log.Fatal(err.Error())
	}
	if err := upload_server.RunUploadServer(); err != nil {
		log.Fatal(err.Error())
	}
}