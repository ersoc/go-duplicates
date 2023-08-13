package main

import (
	"os"
	"time"
)

type Metadata struct {
	Name    string
	Size    int64
	ModTime time.Time
}

func GenerateMetadata(path string) (*Metadata, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	metadata := &Metadata{
		Name:    fileInfo.Name(),
		Size:    fileInfo.Size(),
		ModTime: fileInfo.ModTime(),
	}

	return metadata, nil
}

func main() {

}
