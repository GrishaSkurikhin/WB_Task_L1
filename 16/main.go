package main

import "fmt"

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func partition(data []int, low, high int) int {
	// Опорную точку можно выбрать любую
	pivot := data[low]

	// Сортируем массив: все, что меньше pivot ставим левее, все что больше - правее
	i, j := low, high
	for i < j {
		for data[i] <= pivot && i < high {
			i++;
		}
		for data[j] > pivot && j > low {
			j--
		}

		if i < j {
			data[i], data[j] = data[j], data[i]
		}
	}
	data[low] = data[j]
	data[j] = pivot
	return j
}

func quickSort(data []int, low, high int){
	if low < high {
		p := partition(data, low, high)
		quickSort(data, low, p-1)
		quickSort(data, p+1, high)
	}
}

func sort(data []int) {
	quickSort(data, 0, len(data)-1)
}

func main() {
	a := []int{1, 3, 2, 4, 2}
	sort(a)
	fmt.Println(a)
}
