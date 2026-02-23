package main

import (
	"fmt"
	"strconv"
)

type Row struct {
	km    float64
	price float64
}

type DataSet struct {
	rows []Row
	size int
}

func NewDataSet(data []Row) *DataSet {
	trainingSet := &DataSet{rows: data, size: len(data)}
	return trainingSet
}

func (ds *DataSet) DisplaySataSet() {
	for _, row := range ds.rows {
		fmt.Printf("kilometers: %v, price: %v\n", row.km, row.price)
	}
}

func ConvertDataStringToRow(csvData [][]string) ([]Row, error) {
	dataSetConverted := make([]Row, 0, len(csvData)-1)
	for idx, row := range csvData {
		if idx == 0 {
			continue
		}

		km, err := strconv.ParseFloat(row[0], 64)
		if err != nil {
			return nil, err
		}
		price, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			return nil, err
		}

		dataSetConverted = append(dataSetConverted, Row{km: km, price: price})
	}

	return dataSetConverted, nil
}
