package model

import (
	"log"
	"strconv"
)

type Row struct {
	Km    float64
	Price float64
}

type DataSet struct {
	Rows []Row
	Size int
}

func NewDataSet(data []Row) *DataSet {
	return &DataSet{Rows: data, Size: len(data)}
}

func ConvertDataStringToRow(csvData [][]string) ([]Row, error) {
	dataSetConverted := make([]Row, 0, len(csvData)-1)
	for idx, row := range csvData {
		if idx == 0 {
			continue
		}

		if len(row) != 2 {
			log.Fatal("CSV Wrong Format !!")
		}

		km, err := strconv.ParseFloat(row[0], 64)
		if err != nil {
			return nil, err
		}
		price, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			return nil, err
		}

		dataSetConverted = append(dataSetConverted, Row{Km: km, Price: price})
	}

	return dataSetConverted, nil
}
