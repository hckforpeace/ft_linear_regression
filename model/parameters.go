package model

type Parameters struct {
	Theta0 float64
	Theta1 float64
}

func NewParameters(theta0 float64, theta1 float64) *Parameters {
	return &Parameters{Theta0: theta0, Theta1: theta1}
}

func (p *Parameters) Process(x float64) float64 {
	return p.Theta0 + (p.Theta1 * x)
}
