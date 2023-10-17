package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(2^2+3^2+4^2+...) с использованием конкурентных вычислений.

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	output := 0
	var wg sync.WaitGroup
	var mx sync.Mutex

	for i := range numbers {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			tmp := numbers[i] * numbers[i]
			// Используем mutex для конкурентной записи в ответ
			mx.Lock()
			output += tmp
			mx.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println(output)
}
