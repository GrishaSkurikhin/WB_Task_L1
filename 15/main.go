package main

import (
	"fmt"
	"strings"
)

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }


// Главная проблема приведенного кода - утечка памяти
// Переменная justString имеет указатель на строку v, из-за чего сборщик мусора не удалит v

func createHugeString(strSize int) string {
	// Для построения большой строки используем string.Builder
	var sb strings.Builder
	for i := 0; i < strSize; i++ {
		sb.WriteString("a")
	}
	return sb.String()
}

var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	// Создаем новую строку с помощью string, которая не будет указывать на v
	justString = string(v[:100])
}

func main() {
	someFunc()
	fmt.Println(justString)
}
