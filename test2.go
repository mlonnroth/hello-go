// test2.go
// Testing interfaces
package main

import "fmt"
import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func totalArea(shapes []Shape) (area float64) {
	for _, v := range shapes {
		area += v.Area()
	}
	return area
}

func main() {
	shape1 := Rectangle{10, 20}
	shape2 := Circle{25}
	things := []Shape{shape1, shape2}
	fmt.Println("Interfaces..")
	fmt.Println(things)
	area := totalArea(things)
	fmt.Printf("Total area = %.2f\n", area)
}
