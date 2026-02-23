package main

type Cost struct {
	costs            map[float64]float64
	min              float64
	max              float64
	overAllCostTeta0 float64
	overAllCostTeta1 float64
}


func NewCost(set *DataSet) *Cost {
	cost := &Cost{costs: make(map[float64]float64)}
	if len(set.rows) == 0 {
		return cost
	}
	minKm := set.rows[0].km
	maxKm := set.rows[0].km
	for _, row := range set.rows {
		if row.km < minKm {
			minKm = row.km
		}
		if row.km > maxKm {
			maxKm = row.km
		}
	}
	cost.min = minKm
	cost.max = maxKm
	rangeKm := maxKm - minKm
	for _, row := range set.rows {
		kmNorm := row.km
		if rangeKm != 0 {
			kmNorm = (row.km - minKm) / rangeKm
		}
		cost.costs[kmNorm] = row.price
	}
	return cost
}

func (c *Cost) ComputeCosts(params *Parameters) {
	var predictedPrice float64
	var cost float64
	var overAllCostTeta0 float64
	var overAllCostTeta1 float64

	size := len(c.costs)
	for kmNorm, price := range c.costs {
		predictedPrice = params.Proccess(kmNorm)
		cost = predictedPrice - price
		overAllCostTeta0 += cost
		overAllCostTeta1 += cost * kmNorm
	}

	overAllCostTeta0 = overAllCostTeta0 / float64(size)
	overAllCostTeta1 = overAllCostTeta1 / float64(size)
	c.overAllCostTeta0 = overAllCostTeta0
	c.overAllCostTeta1 = overAllCostTeta1
}

func (c *Cost) Denormalise(params *Parameters) {
	rangeKm := c.max - c.min
	params.teta1 = params.teta1 / rangeKm
	params.teta0 = params.teta0 - (params.teta1 * c.min)
}
