package main

import (
	"encoding/csv"
	"os"

	"gonum.org/v1/plot/plotter"
)

func readData(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func toXYs(set *DataSet) plotter.XYs {
	pts := make(plotter.XYs, len(set.rows))
	for i, row := range set.rows {
		pts[i].X = row.km
		pts[i].Y = row.price
	}
	return pts
}
