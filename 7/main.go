package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.

type syncMap struct {
	mu   sync.RWMutex
	data map[any]any
}

// Конструктор syncMap
func NewSyncMap(size int) syncMap {
	return syncMap{
		data: make(map[any]any, size),
	}
}

// Взятие по ключу из мапы
func (sm *syncMap) Load(key any) (value any, ok bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, ok = sm.data[key]
	return
}

// Положить значение в мапу
func (sm *syncMap) Store(key, value any) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func main() {
	smap := NewSyncMap(5)

	
	numWorkers := 5

	// Кладем значения в мапу конкурентно
	var wg1 sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg1.Add(1)
		go func(id int) {
			defer wg1.Done()
			key := fmt.Sprintf("Key%d", id)
			value := id * 10

			smap.Store(key, value)
		}(i)
	}
	wg1.Wait()

	// Достаем значения из мапы конкурентно
	var wg2 sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg2.Add(1)
		go func(id int) {
			defer wg2.Done()
			key := fmt.Sprintf("Key%d", id)
			if value, ok := smap.Load(key); ok {
				fmt.Printf("%s -> %d\n", key, value)
			} else {
				fmt.Printf("%s not found\n", key)
			}
		}(i)
	}
	wg2.Wait()
}
