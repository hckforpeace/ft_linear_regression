package plot

import (
	"image/color"
	"log"

	"ft_linear_regression/model"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type DrawPlot struct {
	points plotter.XYs
}

func NewDrawPlot(set *model.DataSet) DrawPlot {
	pts := make(plotter.XYs, len(set.Rows))
	for i, row := range set.Rows {
		pts[i].X = row.Km
		pts[i].Y = row.Price
	}
	return DrawPlot{points: pts}
}

func (d *DrawPlot) Save(parameters *model.Parameters, outputFile string) {
	scatter, err := plotter.NewScatter(d.points)
	if err != nil {
		log.Fatal(err)
	}

	p := plot.New()
	p.Title.Text = "Linear Regression"
	p.X.Label.Text = "Km"
	p.Y.Label.Text = "Price"
	p.Add(scatter)

	line := plotter.NewFunction(func(x float64) float64 {
		return parameters.Theta0 + parameters.Theta1*x
	})
	line.Color = color.RGBA{R: 255, A: 255}
	p.Add(line)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, outputFile); err != nil {
		log.Fatal(err)
	}
}
