package tasks

import (
	"fmt"
	"math"
)

/*Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными
параметрами x,y и конструктором.*/

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) Distance(p1 *Point) float64 {
	return math.Sqrt(math.Pow((p.x-p1.x), 2) + math.Pow((p.y-p1.y), 2))
}

func L1_24() {
	p1 := NewPoint(-1, 3)
	p2 := NewPoint(6, 2)
	fmt.Println(p1.Distance(p2))
}
