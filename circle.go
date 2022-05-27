package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (c *Circle) CalcPerimeter() float64 {
	perimeter := 2 * math.Pi * c.Radius
	return perimeter
}

func (c *Circle) CalcArea() float64 {
	area := 2 * c.Radius * c.Radius
	return area
}
