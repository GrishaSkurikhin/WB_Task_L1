package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

// Для определения уникальности символа будем записывать его в мапу
func IsUnique(s string) bool {
	// Сначала приводим строку к нижнему регистру
	lowerString := strings.ToLower(s)
	// Приводим к слайсу рун
	runes := []rune(lowerString)
	chars := make(map[rune]struct{}, len(runes))
	
	for _, char := range runes {
		if _, isExist := chars[char]; isExist {
			return false
		}
		chars[char] = struct{}{}
	} 
	return true
}

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"
	fmt.Printf("%s - %t\n", s1, IsUnique(s1))
	fmt.Printf("%s - %t\n", s2, IsUnique(s2))
	fmt.Printf("%s - %t\n", s3, IsUnique(s3))
}