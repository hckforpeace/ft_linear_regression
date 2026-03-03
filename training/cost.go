package training

import "ft_linear_regression/model"

type normalizeRow struct {
	km    float64
	price float64
}

type Cost struct {
	data             []normalizeRow
	min              float64
	max              float64
	overAllCostTeta0 float64
	overAllCostTeta1 float64
}

func NewCost(set *model.DataSet) *Cost {
	cost := &Cost{data: make([]normalizeRow, 0, len(set.Rows))}
	if len(set.Rows) == 0 {
		return cost
	}
	minKm := set.Rows[0].Km
	maxKm := set.Rows[0].Km
	for _, row := range set.Rows {
		if row.Km < minKm {
			minKm = row.Km
		}
		if row.Km > maxKm {
			maxKm = row.Km
		}
	}
	cost.min = minKm
	cost.max = maxKm
	rangeKm := maxKm - minKm
	for _, row := range set.Rows {
		kmNorm := row.Km
		if rangeKm != 0 {
			kmNorm = (row.Km - minKm) / rangeKm
		}
		cost.data = append(cost.data, normalizeRow{km: kmNorm, price: row.Price})
	}
	return cost
}

func (c *Cost) ComputeCosts(params *model.Parameters) {
	var overAllCostTeta0 float64
	var overAllCostTeta1 float64

	size := len(c.data)
	for _, normalisedRow := range c.data {
		predictedPrice := params.Process(normalisedRow.km)
		cost := predictedPrice - normalisedRow.price
		overAllCostTeta0 += cost
		overAllCostTeta1 += cost * normalisedRow.km
	}

	c.overAllCostTeta0 = overAllCostTeta0 / float64(size)
	c.overAllCostTeta1 = overAllCostTeta1 / float64(size)
}

func (c *Cost) Denormalise(params *model.Parameters) {
	rangeKm := c.max - c.min
	params.Theta1 = params.Theta1 / rangeKm
	params.Theta0 = params.Theta0 - (params.Theta1 * c.min)
}
