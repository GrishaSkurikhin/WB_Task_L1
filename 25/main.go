package main

import (
	"context"
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.


// Реализация с помощью контекста
func sleepContext(t time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()
	select {
	case <-ctx.Done():
		return
	}
}

// Реализация с помощью бесконечного цикла и пакета time
func sleepLoop(t time.Duration) {
	timeStart := time.Now()
	for {
		if time.Since(timeStart) >= t {
			return
		}
	}
}

func main() {
	fmt.Println("sleepContext 2 second")
	sleepContext(2 * time.Second)
	fmt.Println("sleepContext finish")

	fmt.Println()

	fmt.Println("sleepLoop 2 second")
	sleepLoop(2 * time.Second)
	fmt.Println("sleepLoop finish")
}
