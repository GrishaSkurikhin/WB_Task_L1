package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

type point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) point {
	return point{x: x, y: y}
}

func (p *point) Add(p2 point) {
	p.x += p2.x
	p.y += p2.y
}

func (p *point) Mult(p2 point) {
	p.x *= p2.x
	p.y *= p2.y
}

func (p point) Distance(p2 point) float64 {
	return math.Sqrt(math.Pow((p2.x-p.x), 2) + math.Pow((p2.y-p.y), 2))
}

func main() {
	point1 := NewPoint(4.5, 7.6)
	point2 := NewPoint(1.3, 2.0)
	fmt.Println(point1.Distance(point2))
}
