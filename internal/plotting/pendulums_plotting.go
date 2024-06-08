package plotting

import (
	"image/color"

	"github.com/charmbracelet/log"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var width font.Length = vg.Points(0.7)

func PlotAngles(t, phi1, phi2 []float64, maxTime float64, filename string) {
	p := plot.New()

	p.Title.Text = "Coupled Pendulums"
	p.X.Label.Text = "Time (sec)"
	p.Y.Label.Text = "Angle (rad)"

	width = vg.Points(0.7)

	xy1 := make(plotter.XYs, len(t))
	xy2 := make(plotter.XYs, len(t))
	for i := range t {
		xy1[i].X = t[i]
		xy1[i].Y = phi1[i]
		xy2[i].X = t[i]
		xy2[i].Y = phi2[i]
	}

	line1, err := plotter.NewLine(xy1)
	if err != nil {
		log.Fatal(err)
	}

	line1.Width = width
	line1.Color = color.RGBA{R: 255, A: 255}

	line2, err := plotter.NewLine(xy2)
	if err != nil {
		log.Fatal(err)
	}

	line2.Width = width
	line2.Color = color.RGBA{B: 255, A: 255}

	p.Add(line1, line2)
	p.Legend.Add("First pendulum", line1)
	p.Legend.Add("Second pendulum", line2)

	p.X.Max = maxTime / 2
	p.Y.Max = (sliceMax(phi1)+sliceMax(phi2))/2 + (sliceMin(phi1)+sliceMin(phi2))/2

	if err := p.Save(5*vg.Inch, 3*vg.Inch, filename); err != nil {
		log.Fatal(err)
	}
}

func PlotVelocities(t, v1, v2 []float64, maxTime float64, filename string) {
	p := plot.New()

	p.Title.Text = "Velocity of Coupled Pendulums"
	p.X.Label.Text = "Time (sec)"
	p.Y.Label.Text = "Velocity (rad/sec)"

	xy3 := make(plotter.XYs, len(t))
	xy4 := make(plotter.XYs, len(t))
	for i := range t {
		xy3[i].X = t[i]
		xy3[i].Y = v1[i]
		xy4[i].X = t[i]
		xy4[i].Y = v2[i]
	}

	line1, err := plotter.NewLine(xy3)
	if err != nil {
		log.Fatal(err)
	}

	line1.Width = width
	line1.Color = color.RGBA{R: 255, A: 255}

	line2, err := plotter.NewLine(xy4)
	if err != nil {
		log.Fatal(err)
	}

	line2.Width = width
	line2.Color = color.RGBA{B: 255, A: 255}

	p.Add(line1, line2)
	p.Legend.Add("First pendulum", line1)
	p.Legend.Add("Second pendulum", line2)

	p.X.Max = maxTime / 2
	p.Y.Max = (sliceMax(v1)+sliceMax(v2))/2 + (sliceMin(v1)+sliceMin(v2))/2

	if err := p.Save(5*vg.Inch, 3*vg.Inch, filename); err != nil {
		log.Fatal(err)
	}
}

func sliceMin(data []float64) float64 {
	result := data[0]

	for _, value := range data {
		if value < result {
			result = value
		}
	}

	return result
}

func sliceMax(data []float64) float64 {
	result := data[0]

	for _, value := range data {
		if value > result {
			result = value
		}
	}

	return result
}
