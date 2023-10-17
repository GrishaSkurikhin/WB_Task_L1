package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {
	a := 10
	b := 20
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}
