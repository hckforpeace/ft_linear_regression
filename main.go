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

func readProcessCsvtrainingDataSet(filename string) (*DataSet, error) {
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

	teta0, err := strconv.ParseFloat(paramsCsvData[1][0], 64)
	if err != nil {
		return nil, err
	}

	teta1, err := strconv.ParseFloat(paramsCsvData[1][1], 64)
	if err != nil {
		return nil, err
	}

	parameters := NewParameters(teta0, teta1)

	return parameters, nil
}

func main() {
	trainingSet, err := readProcessCsvtrainingDataSet("data.csv")
	if err != nil {
		log.Fatal(err)
		return
	}

	p := plot.New()
	p.Title.Text = "Linear Regression"
	p.X.Label.Text = "Km"
	p.Y.Label.Text = "Price"

	points := plotter.XYs(toXYs(trainingSet))
	scatter, _ := plotter.NewScatter(points)
	p.Add(scatter)

	// trainingSet.DisplaySataSet()

	parameters, err := readProcessCsvParametersDataSet("parameters.csv")
	if err != nil {
		log.Fatal(err)
		return
	}

	cost := NewCost(trainingSet)

	gd := GradientDescent{LearningRate: 0.01, IterationNumber: 1000000}

	gd.gradient_descent(parameters, trainingSet, cost)

	cost.Denormalise(parameters)
	fmt.Printf("teta0: %v, teta1: %v\n", parameters.teta0, parameters.teta1)

	line := plotter.NewFunction(func(x float64) float64 {
		return parameters.teta0 + parameters.teta1*x
	})
	line.Color = color.RGBA{R: 255, A: 255}
	p.Add(line)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, "scatter.png"); err != nil {
		log.Fatal(err)
	}
}
