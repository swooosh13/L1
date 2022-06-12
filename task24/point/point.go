package point

import "math"

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x, y,
	}
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func Distance(a, b *Point) float64 {
	return math.Sqrt((a.GetX() - b.GetX()) * (a.GetX() - b.GetX()) + (a.GetY() - b.GetY())*(a.GetY() - b.GetY()))
}

