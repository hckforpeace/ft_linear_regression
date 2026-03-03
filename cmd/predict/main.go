package main

import (
	"fmt"
	"log"
	"strconv"

	"ft_linear_regression/consts"
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

	fmt.Printf("Predicted price: %v\n", data.Process(value))
}
