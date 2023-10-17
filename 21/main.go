package main

import "fmt"

// Реализовать паттерн «адаптер» на любом примере.

// Клиент обращается с запросом к сервисам
type Client struct {}

func (c Client) PrintTarget(target string) {
	fmt.Println(target)
}

///////////////////////////////////////
type Service interface {
	Target() string
} 

// Сервис 1 генерирует правильную строку
type Service1 struct {}

func (s1 Service1) Target() string {
	return "The default target's behavior."
}

// Сервис 2 генерирует перевернутую строку 
type Service2 struct {}

func (s2 Service2) Target() string {
	return ".roivaheb s'tegrat laicepS"
}

//////////////////////////////////////

// Адаптер для сервиса 2 - переворачивает строку
type Service2Adapter struct {
	s2 Service2
}

func (s2adapter Service2Adapter) Target() string {
	return reverse(s2adapter.s2.Target())
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
} 

func main() {
	client := Client{}
	s1 := Service1{}
	s2 := Service2{}
	fmt.Println("Expected string:")
	client.PrintTarget(s1.Target())
	fmt.Println()

	fmt.Println("Wrong string:")
	client.PrintTarget(s2.Target())
	fmt.Println()

	fmt.Println("Wrong string after adapter:")
	s2Adapter := Service2Adapter{s2}
	client.PrintTarget(s2Adapter.Target())
}