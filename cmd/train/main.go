package main

import (
	"fmt"
	"log"

	"ft_linear_regression/consts"
	"ft_linear_regression/csvio"
	"ft_linear_regression/plot"
	"ft_linear_regression/training"

)


func main() {
	trainingSet, err := csvio.ReadTrainingDataSet(consts.TrainingDataFile)
	if err != nil {
		log.Fatal(err)
	}

	drawPlot := plot.NewDrawPlot(trainingSet)

	cost := training.NewCost(trainingSet)

	parameters, err := csvio.ReadParameters(consts.ParametersFile)
	if err != nil {
		log.Fatal(err)
	}

	gd := training.GradientDescent{LearningRate: 0.01, IterationNumber: 1000000}
	gd.Run(parameters, cost)

	cost.Denormalise(parameters)
	fmt.Printf("theta0: %v, theta1: %v\n", parameters.Theta0, parameters.Theta1)

	drawPlot.Save(parameters, consts.OutputPlotFile)

	if err := csvio.WriteParameters(consts.ParametersFile, parameters); err != nil {
		log.Fatal(err)
	}
}
