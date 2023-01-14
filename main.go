package main

import (
	"fmt"
	"reflect"
)

type shape interface {
	area() float64
}
type square struct {
	length float64
}

type triangle struct {
	base   float64
	height float64
}

func main() {
	sq := square{5.25}
	tri := triangle{base: 1.00, height: 2.00}

	printArea(sq)
	printArea(tri)
}

func (tr triangle) area() float64 {
	return 0.5 * tr.base * tr.height
}
func (sq square) area() float64 {
	return sq.length * sq.length
}

func printArea(sh shape) {
	fmt.Println(reflect.TypeOf(sh), sh.area())
}
