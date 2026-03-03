package main 

import(
	"fmt"
	"log"
	"math"
	"ft_linear_regression/csvio"
	"ft_linear_regression/consts"
)

func main() {
	trainingSet, err := csvio.ReadTrainingDataSet(consts.TrainingDataFile)
	if err != nil {
		log.Fatal(err)
	}
	parameters, err := csvio.ReadParameters(consts.ParametersFile)
	if err != nil {
		log.Fatal(err)
	}

	mse := trainingSet.ComputeMSE(parameters)
	fmt.Printf("Mean Squared Error: %v\n", mse)

	minPrice, maxPrice := trainingSet.Rows[0].Price, trainingSet.Rows[0].Price
	for _, row := range trainingSet.Rows {
		if row.Price < minPrice {
			minPrice = row.Price
		}
		if row.Price > maxPrice {
			maxPrice = row.Price
		}
	}
	precision := math.Sqrt(mse) / (maxPrice - minPrice) * 100
	fmt.Printf("Precision: %.2f%% error over price range\n", precision)
}
