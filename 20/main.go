package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func reverseWords(s string) string {
	// Разделяем строку на отдельные слова
	words := strings.Fields(s)

	// Разворачиваем слова
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Объеднияем слова обратно в строку
	result := strings.Join(words, " ")

	return result
}

func main() {
	s := "snow dog sun"
	fmt.Println(reverseWords(s))
}
