package filesystem



import (
	"io"
	"log"
	"os"
)



type FileSystem struct {
	rootPath string
	logFile  *os.File
}



func NewFileSystem() *FileSystem {
	file_system := new(FileSystem)
	return file_system
}


func (file_system *FileSystem) InitFileSystem(rootPath string) error {
	// Create root path.
	err := file_system.CreateDirIfNotExists(rootPath)
	if err != nil {
		return err
	}
	file_system.rootPath = rootPath

	// Creates the general upload-server log file.
	logFile, err := os.Create(file_system.rootPath + "/daemon.log")
	if err != nil {
		return err
	}

	// Redirects log for Stdout and the created log file.
	writer := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(writer)
	file_system.logFile = logFile

	return nil
}


func (file_system *FileSystem) CreateDirIfNotExists(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(path, os.ModePerm)
		} else {
			return err
		}
	}
	return nil
}