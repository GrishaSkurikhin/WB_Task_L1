package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.

func BinarySearch(a []int, target int) (int, bool) { 
	l, r := 0, len(a)-1 
	for l <= r { 
		m := (l + r) / 2 
		if a[m] > target { 
			r = m - 1 
		} else if a[m] < target { 
			l = m + 1 
		} else { 
			return m, true 
		} 
	} 
	return 0, false 
}

func main() {
	a := []int{1, 2, 4, 8, 9, 16}
	target := 4
	index, isExist := BinarySearch(a, target)
	fmt.Println(index, isExist)
}