package filesystem

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

// FileSystem - Filesystem data structure
type FileSystem struct {
	rootPath string
	logFile  *os.File
	indexes  map[string]string
}

func newFileSystem(rootPath string) *FileSystem {
	return &FileSystem{rootPath, nil, map[string]string{}}
}

func (fileSystem *FileSystem) init() error {
	// Create root path.
	err := fileSystem.createDirIfNotExists(fileSystem.rootPath)
	if err != nil {
		return err
	}

	// Creates the general upload-server log file.
	logPath := fmt.Sprintf("%s/%d -> %s.log", fileSystem.rootPath, os.Getpid(), time.Now().String())
	logFile, err := os.Create(logPath)
	if err != nil {
		return err
	}

	// Redirects log for Stdout and the created log file.
	writer := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(writer)
	fileSystem.logFile = logFile

	return nil
}

func (fileSystem *FileSystem) stop() {
	if err := fileSystem.logFile.Close(); err != nil {
		log.Fatal(err.Error())
	}
}

func (fileSystem *FileSystem) createDirIfNotExists(path string) error {
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

func (fileSystem *FileSystem) copyFile(originalFile io.Reader, fileName string) error {
	fileUUID := uuid.New().String()
	path := filepath.Join("data/", fileUUID)
	createdFile, err := os.Create(path)
	if err != nil {
		return err
	}

	createdFile, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}

	_, err = io.Copy(createdFile, originalFile)
	if err != nil {
		return err
	}

	fileSystem.indexes[fileName] = fileUUID

	return nil
}

func (fileSystem *FileSystem) getIndexes() map[string]string {
	return fileSystem.indexes
}
