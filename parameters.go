package main

import "fmt"

type Parameters struct {
	teta0 float64
	teta1 float64
}

func NewParameters(teta0 float64, teta1 float64) *Parameters {
	return &Parameters{teta0: teta0, teta1: teta1}
}

func (p *Parameters) displayParams() {
	fmt.Printf("teta0: %v, teta1: %v\n", p.teta0, p.teta1)
}

func (p *Parameters) Proccess(x float64) float64 {
	output := p.teta0 + (p.teta1 * x)

	return output
}
