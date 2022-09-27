package main

type Shape interface {
	area() float32
}

type Circle struct {
	x, y, radius int
}

func (circle Circle) area() float32 {
	return float32(circle.radius)
}

type Ractangle struct {
	x, y int
}

func main() {

}
