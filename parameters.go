package main


type Parameters struct {
	theta0 float64
	theta1 float64
}

func NewParameters(teta0 float64, teta1 float64) *Parameters {
	return &Parameters{theta0: teta0, theta1: teta1}
}

func (p *Parameters) Process(x float64) float64 {
	output := p.theta0 + (p.theta1 * x)

	return output
}
