package main

import (
	"fmt"
	"strconv"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func setBit(num int64, i int, bitValue int) int64 {
	switch bitValue {
	case 1:
		// Побитовое ИЛИ с маской вида 00010000
		return num | (int64(1) << i)
	case 0:
		// Побитовое И с маской вида 11101111
		return num & ^(int64(1) << i)
	default:
		return num
	}
}

// Вывод числа в бинарном виде
func printBinary(num int64) {
	binString := strconv.FormatInt(num, 2)
	fmt.Println(binString)
}

func main() {
	var num int64
	num = 8
	printBinary(num)
	num = setBit(num, 2, 1)
	printBinary(num)
}
