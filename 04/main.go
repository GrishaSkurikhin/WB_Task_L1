package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

// Функция worker, совершающая полезные действия
func worker(id int, dataChan <-chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// Ждем закрытия канала данных и завершаем функцию
		if data, ok := <-dataChan; !ok {
			fmt.Printf("Worker %d finished work\n", id)
			return
		} else {
			fmt.Printf("Worker %d, data: %v\n", id, data)
		}
	}
}

func main() {
	numWorkers := 5
	dataChan := make(chan interface{}, numWorkers)

	var wg sync.WaitGroup
	// Запускаем воркеров в отдельных горутинах
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChan, &wg)
	}

	// Пишем некоторые данные в канал
	go func() {
		for {
			dataChan <- rand.Intn(50)
			time.Sleep(1*time.Second)
		}
	}()
	
	// Запись в канал и обработка воркерами данных выполняется до сигнала на завершение программы
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-sigCh
	
	// Закрываем канал и ждем завершения работы всех горутин
	close(dataChan)
	wg.Wait()

	fmt.Println("Program completed")
}
