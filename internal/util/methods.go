package util

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateMiniUrl() string {
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

//func CheckMapUrl(mapURL map[string]string, path string) string {
//	return mapURL[path]
//}
//
//func SetMapUrl(mapURL map[string]string, path string) (string, error) {
//	// добавляет урл в словарь mapURL
//	for miniURL, URL := range mapURL {
//		if URL == path {
//			return miniURL, errors.New("данный URL уже записан")
//		}
//	}
//	for {
//		miniURL := GenerateMiniUrl()
//		if mapURL[miniURL] != "" {
//			continue
//		}
//		mapURL[miniURL] = path
//		return miniURL, nil
//	}
//
//}
