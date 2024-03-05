package structsmethodsandinterfaces

type Triangle struct {
	Base   float64
	Height float64
}

var _ Shape = Triangle{}

func (r Triangle) Area() float64 {
	return r.Base * r.Height / 2.0
}

func (r Triangle) Perimeter() float64 {
	return 0.0
}
