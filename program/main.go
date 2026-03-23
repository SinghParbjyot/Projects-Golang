package main

import "fmt"

type triangle struct {
	heigth float64
	base   float64
}
type square struct {
	sideLength float64
}
type shape interface {
	getArea() float64
}

func main() {
	s := square{sideLength: 10}
	t := triangle{heigth: 6, base: 4}
	printArea(s)
	printArea(t)
}
func printArea(s shape) {
	fmt.Println(s.getArea())
}
func (t triangle) getArea() float64 {
	return (t.base * t.heigth) / 2
}
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
