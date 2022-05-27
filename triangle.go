package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t *Triangle) CalcPerimeter() float64 {
	perimeter := t.Side * 3
	return perimeter
}

func (t *Triangle) CalcArea() float64 {
	area := (math.Sqrt(3) / 4) * (t.Side * t.Side)
	return area
}
