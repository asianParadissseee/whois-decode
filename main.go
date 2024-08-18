package main

import (
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io"
	"bytes"
)

// decodeMessedUpString декодирует строку, которая была неправильно закодирована
func decodeMessedUpString(s string) (string, error) {
	// Преобразуем строку в байтовый срез
	byteArray := []byte(s)

	// Преобразуем каждый байт, используя Windows-1251 кодировку
	decoder := charmap.Windows1251.NewDecoder()
	decodedStr, err := io.ReadAll(transform.NewReader(bytes.NewReader(byteArray), decoder))
	if err != nil {
		return "", err
	}

	return string(decodedStr), nil
}

func main() {
	// Пример неправильно закодированной строки (UTF-8, интерпретированная как Windows-1251)
	messedUpString := "РђС…РјР°С‚РѕРІ Р Р°СѓС„" // Пример строки

	// Декодируем строку
	correctString, err := decodeMessedUpString(messedUpString)
	if err != nil {
		fmt.Println("Ошибка декодирования:", err)
	} else {
		fmt.Println(correctString) // Ожидаемый вывод: Ахматов Рауф
	}
}
