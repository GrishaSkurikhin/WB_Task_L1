package main

import (
	"fmt"
	"math"
)

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов.
// Последовательность в подмножноствах не важна.

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// Используем мапу для создания групп
	// Ключ - десятичный разряд, значение - слайс температур в группе 
	temperatureGroups := make(map[float64][]float64, len(temperatures))

	for _, temp := range temperatures {
		// Определяем десятичный рязряд числа 
		groupKey := math.Trunc(temp/10) * 10
		// Добавляем в группу
		temperatureGroups[groupKey] = append(temperatureGroups[groupKey], temp)
	}

	for key, value := range temperatureGroups {
		fmt.Printf("Group %g: %v\n", key, value)
	}
}
