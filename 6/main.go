package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

// Выполнение горутины завершается по выходу из функции
func cancelWithReturn() {
	// Do something ...
	fmt.Println("Gourtine cancelWithReturn finished work")
	return
}

// Горутина ждет завершения других запущенных горутин
func cancelWithWaitGroup(wg *sync.WaitGroup) {
	// Do something ...
	wg.Wait()
	fmt.Println("Gourtine cancelWithWaitGroup finished work")
	return
}

// Выполнение горутины прекращается по закрытию переданного канала
func cancelWithChannel(doneChan <-chan struct{}) {
	for {
		select {
		case <- doneChan:
			fmt.Println("Gourtine cancelWithChannel finished work")
			return
		default:
			// Do something ...
		}
	}
}

// Выполнение горутины прекращается по завершению переданного контекста
func cancelWithContext(ctx context.Context) {
	for {
		select {
		case <- ctx.Done():
			fmt.Println("Gourtine cancelWithContext finished work")
			return
		default:
			// Do something ...
		}
	}
}



func main() {
	go cancelWithReturn()

	var wg sync.WaitGroup
	wg.Add(1)
	go cancelWithWaitGroup(&wg)
	wg.Done()

	doneChan := make(chan struct{})
	go cancelWithChannel(doneChan)
	doneChan <- struct{}{}

	ctx, cancel := context.WithCancel(context.Background())
	go cancelWithContext(ctx)
	cancel()

	time.Sleep(time.Second)
}