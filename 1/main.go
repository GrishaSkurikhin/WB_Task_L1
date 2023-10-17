package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

// Дочерняя структура
type Human struct {
    Name string
    Age  int
}

// Метод дочерней структуры
func (h *Human) Speak() {
    fmt.Printf(`%s says: "Hello, I'm %d years old"`+"\n", h.Name, h.Age)
}

// Родительская структура
type Action struct {
    Human
    ActionName string
}

// Метод родительской структуры
func (a *Action) Do() {
    fmt.Printf("%s %s\n", a.Name, a.ActionName)
}

func main() {
    a := Action{
        Human:      Human{Name: "Jhon", Age: 30},
        ActionName: "run",
    }

    // Вывоз метода дочерней структуры из родительской
    a.Speak()
    // Вызов метода родительской структуры
    a.Do()
}