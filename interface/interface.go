package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}

type Cycle struct {
	Radius float32
}
type Rectangle struct {
	width float32
	heigh float32
}

func (c Cycle) Area() float32 {
	return math.Pi * c.Radius
}
func (r Rectangle) Area() float32 {
	return r.heigh * r.width
}

func main() {
	fmt.Println("Demo interface in golang")
	var shape Shape = &Cycle{Radius: 10}
	fmt.Println("Cycle:     ", shape.Area())
	shape = &Rectangle{width: 10, heigh: 20}
	fmt.Println("Rectangle: ", shape.Area())
}
