package storage

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type StorageMapURL struct {
	// Словарь для хванения URL
	MapURL map[string]string
}

type Storage interface {
	// Интерфейс для работы с хранилищем
	Read(miniURL string) string
	Record(bigURL string) (string, error)
	FullRecord(bigURL, miniURL string) error
}

func ReadStorage(s Storage, miniURL string) string {
	return s.Read(miniURL)
}

func RecordStorage(s Storage, bigURL string) (string, error) {
	return s.Record(bigURL)
}

func (s *StorageMapURL) Read(miniURL string) string {
	return s.MapURL[miniURL]
}

func (s *StorageMapURL) Record(bigURL string) (string, error) {
	for miniURL, URL := range s.MapURL {
		if URL == bigURL {
			return miniURL, errors.New("данный URL уже записан")
		}
	}
	for {
		miniURL := GenerateMiniURL()
		if s.MapURL[miniURL] != "" {
			continue
		}
		s.MapURL[miniURL] = bigURL
		return miniURL, nil
	}
}

func (s *StorageMapURL) FullRecord(miniURL, bigURL string) error {

	if s.MapURL[miniURL] != "" {
		return errors.New("запись уже существует")
	}
	s.MapURL[miniURL] = bigURL
	return nil

}

var _ FileWorkInterface = &FileWorkStruct{}

type URLStorageStruct struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type FileWorkStruct struct {
	File *os.File
	Path string
}

func (f *FileWorkStruct) Close() {
	f.File.Close()
}

func (f *FileWorkStruct) OpenFileRead(path string) {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Println(err)
	}
	f.File = file
	f.Path = path
}

func (f *FileWorkStruct) OpenFileWrite(path string) {

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Println(err)
	}
	f.File = file
	f.Path = path
}

func (f FileWorkStruct) ReadFile() string {
	fileData, err := os.ReadFile(f.File.Name())
	if err != nil {
		log.Println(err)
	}
	return string(fileData)
}

func (f *FileWorkStruct) WriteFile(str URLStorageStruct) {
	data, err := json.Marshal(str)
	if err != nil {
		log.Println(err)
	}
	_, err = f.File.Write([]byte(string(data) + "\n"))
	if err != nil {
		log.Println(err)
	}

}

type FileWorkInterface interface {
	OpenFileRead(path string)
	OpenFileWrite(path string)
	ReadFile() string
	WriteFile(str URLStorageStruct)
	Close()
}

func GenerateMiniURL() string {
	// Функция, которая генерирует сокращенный путь
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 6
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String() // Например "4eaxo3"
	return str
}
