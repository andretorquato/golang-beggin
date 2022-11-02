package geometry

type Point struct {
	X, Y float64
}

func (p Point) Sum() float64 {
	return p.X + p.Y
}

func (p Point) Sub() float64 {
	return p.X - p.Y
}
