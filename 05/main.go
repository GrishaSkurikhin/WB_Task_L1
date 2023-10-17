package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

const (
	N = 3
)

func sender(ch chan<- int, ctx context.Context) {
	defer close(ch)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			ch <- rand.Intn(50)
			time.Sleep(time.Second/2)
		}
	}
}

func main() {
	ch := make(chan int)
	// Будем завершать работу горутин с помощью контекста
	ctx, cancel := context.WithTimeout(context.Background(), N*time.Second)
	defer cancel()

	// Запускаем отправителя, передаем ему созданый контекст
	go sender(ch, ctx)

	// Читаем данные из канала до завершения контекста
	for {
		select {
		case <-ctx.Done():
			return
		case data, ok := <-ch:
            if !ok {
                return
            } else {
				fmt.Println(data)
			}
		}
	}
}
