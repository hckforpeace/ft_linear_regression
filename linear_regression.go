package main

import (
	"fmt"
	"math"
)

const Epsilon = 1e-9

type GradientDescent struct {
	LearningRate    float64
	IterationNumber int
}

func (gd *GradientDescent) gradientDescent(params *Parameters, set *DataSet, cost *Cost) {
	var descentTeta1 float64
	var descentTeta0 float64

	for i := 0; i < gd.IterationNumber; i++ {
		cost.ComputeCosts(params)
		descentTeta0 = gd.LearningRate * cost.overAllCostTeta0
		descentTeta1 = gd.LearningRate * cost.overAllCostTeta1
		params.theta0 -= descentTeta0
		params.theta1 -= descentTeta1

		if math.Abs(descentTeta0) < Epsilon && math.Abs(descentTeta1) < Epsilon {
			fmt.Printf("converged at iteration %d\n", i)
			return
		}
	}

	fmt.Printf("gradient done after %d iterations (max reached)\n", gd.IterationNumber)
}
