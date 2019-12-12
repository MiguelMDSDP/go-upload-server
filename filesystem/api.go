package filesystem

import (
	"io"
	"log"
)

var instance *FileSystem

func init() {
	if instance == nil {
		instance = newFileSystem("data")
	}
	err := instance.init()
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Stop stalls the FileSystem routine
func Stop() {
	instance.stop()
}

// CopyFile copies the incoming file to destiny directory with an UUID
func CopyFile(originalFile io.Reader, fileName string) error {
	return instance.copyFile(originalFile, fileName)
}

// Indexes return the filesystem indexes
func Indexes() map[string]string {
	return instance.getIndexes()
}
