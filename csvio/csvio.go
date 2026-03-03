package csvio

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"

	"ft_linear_regression/model"
)

func ReadData(filename string) ([][]string, error) {
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

func ReadTrainingDataSet(filename string) (*model.DataSet, error) {
	csvData, err := ReadData(filename)
	if err != nil {
		return nil, err
	}

	rows, err := model.ConvertDataStringToRow(csvData)
	if err != nil {
		return nil, err
	}
	return model.NewDataSet(rows), nil
}

func ReadParameters(filename string) (*model.Parameters, error) {
	csvData, err := ReadData(filename)
	if err != nil {
		return nil, err
	}

	if len(csvData) != 2 {
		return nil, errors.New("parameters CSV: expected 2 rows (header + data)")
	}
	if len(csvData[1]) != 2 {
		return nil, errors.New("parameters CSV: expected 2 columns (theta0, theta1)")
	}

	theta0, err := strconv.ParseFloat(csvData[1][0], 64)
	if err != nil {
		return nil, err
	}

	theta1, err := strconv.ParseFloat(csvData[1][1], 64)
	if err != nil {
		return nil, err
	}

	return model.NewParameters(theta0, theta1), nil
}

func WriteParameters(filename string, params *model.Parameters) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	if err := writer.Write([]string{"theta0", "theta1"}); err != nil {
		return fmt.Errorf("writing CSV header: %w", err)
	}

	theta0 := strconv.FormatFloat(params.Theta0, 'f', 2, 64)
	theta1 := strconv.FormatFloat(params.Theta1, 'f', 2, 64)
	if err := writer.Write([]string{theta0, theta1}); err != nil {
		return fmt.Errorf("writing CSV record: %w", err)
	}

	writer.Flush()
	return writer.Error()
}
