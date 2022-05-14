package util

import (
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
