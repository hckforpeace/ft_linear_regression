package training

import (
	"fmt"
	"math"

	"ft_linear_regression/model"
)

const Epsilon = 1e-9

type GradientDescent struct {
	LearningRate    float64
	IterationNumber int
}

func (gd *GradientDescent) Run(params *model.Parameters, cost *Cost) {
	for i := 0; i < gd.IterationNumber; i++ {
		cost.ComputeCosts(params)
		descentTeta0 := gd.LearningRate * cost.overAllCostTeta0
		descentTeta1 := gd.LearningRate * cost.overAllCostTeta1
		params.Theta0 -= descentTeta0
		params.Theta1 -= descentTeta1

		if math.Abs(descentTeta0) < Epsilon && math.Abs(descentTeta1) < Epsilon {
			fmt.Printf("converged at iteration %d\n", i)
			return
		}
	}

	fmt.Printf("gradient done after %d iterations (max reached)\n", gd.IterationNumber)
}
