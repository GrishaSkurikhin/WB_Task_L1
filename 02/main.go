package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает
// значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	output := make([]int, len(numbers))

	// WaitGroup считает запущенные горутины, а потом ждет их завершения
	var wg sync.WaitGroup

	for i := range numbers {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			//Гонки не возникает, т.к. каждая горутина работает только с одним индексом 
			output[i] = numbers[i] * numbers[i]
		}(i)
	}

	wg.Wait()

	for _, num := range output {
        fmt.Println(num)
    }
}
