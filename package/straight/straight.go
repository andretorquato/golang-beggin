package main

import "math"

// Point of on line
type Point struct {
	x float64
	y float64
}

func cateto(a, b Point) (cx, cy float64) {
	cx = math.Abs(a.x - b.x)
	cy = math.Abs(a.y - b.y)
	return
}

func Distance(a, b Point) float64 {
	cx, cy := cateto(a, b)
	return math.Sqrt(math.Pow(cx, 2)) + math.Pow(cy, 2)
}
