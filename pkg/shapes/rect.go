package shapes

type Rectangle struct {
	X, Y float64
}

var _ Shape = (*Rectangle)(nil)

func (r *Rectangle) Area() float64 {
	return r.X * r.Y
}

func (r *Rectangle) Perimeter() float64 {
	return (r.X + r.Y) * 2
}
