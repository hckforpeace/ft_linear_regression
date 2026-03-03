package main 

import (
	"fmt"
	"log"

	"ft_linear_regression/consts"
	"strconv"
	"ft_linear_regression/csvio"
)

func main() {
	data, err := csvio.ReadParameters(consts.ParametersFile)
	if err != nil {
		log.Fatal(err)
	}
	var km string
	fmt.Printf("Input a kilometrage: ")
	fmt.Scan(&km)
	
	value, err := strconv.ParseFloat(km, 64)
	if err != nil {
		log.Fatal(err)
	}
	

	price := data.Theta0 + data.Theta1 *  value
	fmt.Printf("Predicted price: %v\n", price)
}
