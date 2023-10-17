package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

// Структура-счетчик, использующая atomic операции
type atomicCounter struct {
	value int64
}

func NewAtomicCounter() atomicCounter {
	return atomicCounter{value: 0}
}

func (c *atomicCounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *atomicCounter) GetValue() int64 {
	return atomic.LoadInt64(&c.value)
}

// Структура-счетчик, использующий mutex
type mutexCounter struct {
	mu sync.RWMutex
	value int64
}

func NewMutexCounter() mutexCounter {
	return mutexCounter{value: 0}
}

func (c *mutexCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++ 
}

func (c *mutexCounter) GetValue() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup

	numGoroutines := 10

	aCounter := NewAtomicCounter()
	mCounter := NewMutexCounter()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			aCounter.Increment()
			mCounter.Increment()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Atomic counter: ", aCounter.GetValue())
	fmt.Println("Mutex counter: ", mCounter.GetValue())
}
