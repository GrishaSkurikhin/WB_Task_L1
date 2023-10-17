package main

import "fmt"

// Разработать программу, которая в рантайме способна
// определить тип переменной: int, string, bool, channel из переменной типа interface{}.

func determineType(value interface{}) {
	// Для определения типа переменной используется оператор switch
	switch v := value.(type) {
	case int:
		fmt.Printf("Int, value: %v\n", v)
	case string:
		fmt.Printf("String, value: %v\n", v)
	case bool:
		fmt.Printf("Bool, value: %v\n", v)
	case chan int:
		fmt.Printf("Channel, value: %v\n", v)
	default:
		fmt.Printf("Unknown type\n")
	}
}

func main() {
	i := 10
	determineType(i)

	s := "asd"
	determineType(s)

	b := false
	determineType(b)

	c := make(chan int)
	determineType(c)
	
	u := 5.4
	determineType(u)
}
