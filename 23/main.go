package main

import "fmt"

// Удалить i-ый элемент из слайса.

func remove(slice []int, s int) []int {
	// Делаем срез до указаного индекса и после него, а затем объединяем
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice = remove(slice, 2)
	fmt.Println(slice)
}
