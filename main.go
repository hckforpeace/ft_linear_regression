package main

import (
	"errors"
	"fmt"
	"image/color"
	"log"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const (
	TrainingDataFile = "data.csv"
	ParametersFile   = "parameters.csv"
	OutputPlotFile   = "scatter.png"
)

func readProcessCsvTrainingDataSet(filename string) (*DataSet, error) {
	trainingCsvData, err := readData(filename)
	if err != nil {
		return nil, err
	}

	trainingProcessedData, err := ConvertDataStringToRow(trainingCsvData)
	if err != nil {
		return nil, err
	}
	trainingSet := NewDataSet(trainingProcessedData)
	return trainingSet, nil
}

func readProcessCsvParametersDataSet(filename string) (*Parameters, error) {
	paramsCsvData, err := readData(filename)
	if err != nil {
		return nil, err
	}

	if len(paramsCsvData) != 2 {
		return nil, errors.New("Error: wrong format")
	}
	if len(paramsCsvData[1]) != 2 {
		return nil, errors.New("Error: wrong format")
	}

	theta0, err := strconv.ParseFloat(paramsCsvData[1][0], 64)
	if err != nil {
		return nil, err
	}

	theta1, err := strconv.ParseFloat(paramsCsvData[1][1], 64)
	if err != nil {
		return nil, err
	}

	parameters := NewParameters(theta0, theta1)

	return parameters, nil
}

func main() {
	trainingSet, err := readProcessCsvTrainingDataSet(TrainingDataFile)
	if err != nil {
		log.Fatal(err)
	}

	p := plot.New()
	p.Title.Text = "Linear Regression"
	p.X.Label.Text = "Km"
	p.Y.Label.Text = "Price"

	points := plotter.XYs(toXYs(trainingSet))
	scatter, err := plotter.NewScatter(points)
	if err != nil {
		log.Fatal(err)
	}
	p.Add(scatter)


	parameters, err := readProcessCsvParametersDataSet(ParametersFile)
	if err != nil {
		log.Fatal(err)
	}

	cost := NewCost(trainingSet)

	gd := GradientDescent{LearningRate: 0.01, IterationNumber: 1000000}

	gd.gradientDescent(parameters, trainingSet, cost)

	cost.Denormalise(parameters)
	fmt.Printf("teta0: %v, teta1: %v\n", parameters.theta0, parameters.theta1)

	line := plotter.NewFunction(func(x float64) float64 {
		return parameters.theta0 + parameters.theta1*x
	})
	line.Color = color.RGBA{R: 255, A: 255}
	p.Add(line)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, OutputPlotFile); err != nil {
		log.Fatal(err)
	}
}
