package util

import (
	"encoding/json"
	"fmt"
	"github.com/ZhuravlevDmi/serviceURL/internal/config"
	"github.com/ZhuravlevDmi/serviceURL/internal/storage"
	"log"
	"strings"
)

func UpdateStorageMapURL(mapURL storage.Storage, info string) {
	// функция принимает info(это информация с файла, там лежат json структуры URLStorageStruct)
	// и это функция записывает эту инфу в mapURL
	jsonListURL := strings.Split(string(info), "\n")
	for _, d := range jsonListURL {

		if d == "" {
			continue
		}

		var j storage.URLStorageStruct
		err := json.Unmarshal([]byte(d), &j)
		if err != nil {
			fmt.Println(err)
			continue
		}

		e := mapURL.FullRecord(j.ID, j.URL)
		if e != nil {
			log.Println(e)
		}

	}
}

func CheckFile(cfgAdr config.ConfigAdress,
	f storage.FileWorkStruct,
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
