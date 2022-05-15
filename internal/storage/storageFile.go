package storage

import (
	"encoding/json"
	"log"
	"os"
)

var _ FileWorkInterface = &FileWorkStruct{}

type URLStorageStruct struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type FileWorkStruct struct {
	file *os.File
}

func (f *FileWorkStruct) Close() {
	f.file.Close()
}

func (f *FileWorkStruct) OpenFileRead(path string) {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Println(err)
	}
	f.file = file
}

func (f *FileWorkStruct) OpenFileWrite(path string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Println(err)
	}
	f.file = file
}

func (f FileWorkStruct) ReadFile() string {
	fileData, _ := os.ReadFile(f.file.Name())
	return string(fileData)
}

func (f FileWorkStruct) WriteFile(str URLStorageStruct) {
	data, err := json.Marshal(str)
	if err != nil {
		log.Println(err)
	}
	f.file.Write([]byte(string(data) + "\n"))
}

type FileWorkInterface interface {
	OpenFileRead(path string)
	OpenFileWrite(path string)
	ReadFile() string
	WriteFile(str URLStorageStruct)
	Close()
}
