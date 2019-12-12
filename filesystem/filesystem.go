package filesystem

import (
	"io"
	"log"
	"os"
)

// FileSystem - Filesystem data structure
type FileSystem struct {
	rootPath string
	logFile  *os.File
}

// NewFileSystem creates new FileSystem objects
func NewFileSystem() *FileSystem {
	fileSystem := new(FileSystem)
	return fileSystem
}

// Init initiates the FileSystem
func (fileSystem *FileSystem) Init(rootPath string) error {
	// Create root path.
	err := fileSystem.CreateDirIfNotExists(rootPath)
	if err != nil {
		return err
	}
	fileSystem.rootPath = rootPath

	// Creates the general upload-server log file.
	logFile, err := os.Create(fileSystem.rootPath + "/upload-server.log")
	if err != nil {
		return err
	}

	// Redirects log for Stdout and the created log file.
	writer := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(writer)
	fileSystem.logFile = logFile

	return nil
}

// Stop destroys the FileSystem
func (fileSystem *FileSystem) Stop() {
	if err := fileSystem.logFile.Close(); err != nil {
		log.Fatal(err.Error())
	}
}

// CreateDirIfNotExists creates the dir "path" if it doesn't exists
func (fileSystem *FileSystem) CreateDirIfNotExists(path string) error {
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
