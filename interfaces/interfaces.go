package main

import "C"
import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

// Rect
type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height		
}

// Circle
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2*math.Pi * c.radius
}

func measure(g geometry)  {
	fmt.Printf("Type: %T\n", g)
	fmt.Printf("Area: %.2f\n", g.area())
	fmt.Printf("Perimeter: %.2f\n", g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 10}

	measure(r)
	measure(c)
}


