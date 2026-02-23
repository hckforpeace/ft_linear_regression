package main

import "fmt"

type GradientDescent struct {
	LearningRate    float64
	IterationNumber int
}

func (gd *GradientDescent) gradientDescent(params *Parameters, set *DataSet, cost *Cost) {
	var descentTeta1 float64
	var descentTeta0 float64

	for i := 0; i < gd.IterationNumber; i++ {
		cost.ComputeCosts(params)
		descentTeta0 = float64(gd.LearningRate) * cost.overAllCostTeta0
		descentTeta1 = float64(gd.LearningRate) * cost.overAllCostTeta1
		params.theta0 -= descentTeta0
		params.theta1 -= descentTeta1
	}

	fmt.Printf("gradient done\n")
}
