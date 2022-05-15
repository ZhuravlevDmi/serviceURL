package util

import (
	"encoding/json"
	"fmt"
	"github.com/ZhuravlevDmi/serviceURL/internal/config"
	"github.com/ZhuravlevDmi/serviceURL/internal/storage"
	"log"
	"math/rand"
	"strings"
	"time"
)

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

func UpdateStorageMapURL(mapURL storage.Storage, info string) {
	// функция принимает info(это информация с файла, там лежат json структуры URLStorageStruct)
	// и это функция записывает эту инфу в mapURL
	jsonListURL := strings.Split(string(info), "\n")
	for _, d := range jsonListURL {
		var j storage.URLStorageStruct
		err := json.Unmarshal([]byte(d), &j)
		if err != nil {
			fmt.Println(err)
		}

		e := mapURL.FullRecord(j.ID, j.URL)
		if e != nil {
			log.Println(e)
		}

	}
}

func CheckFile(cfgAdr config.ConfigAdress,
	f storage.FileWorkInterface,
	mapURL storage.Storage) {
	if cfgAdr.PATHFile == "" {
		return
	}
	path := cfgAdr.PATHFile
	f.OpenFileRead(path)
	defer f.Close()
	dataFile := f.ReadFile()

	if dataFile == "" {
		return
	}

	UpdateStorageMapURL(mapURL, dataFile)

}
