package main

import (
	"fmt"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func main() {
	// Первый канал
	inputChan := make(chan int)
	// Второй канал
	resultChan := make(chan int)

	// Пишем данные в первый канал
	go func() {
		defer close(inputChan)
		numbers := []int{1, 2, 3, 4, 5}
		for _, num := range numbers {
			inputChan <- num
		}
	}()

	// Достаем данные из первого канал и пишем во второй
	go func() {
		defer close(resultChan)
		for num := range inputChan {
			resultChan <- num * 2
		}
	}()
	
	// Выводим данные из второго канала
	for result := range resultChan {
		fmt.Println(result)
	}
}
