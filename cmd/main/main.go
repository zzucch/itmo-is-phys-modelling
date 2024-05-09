package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"math"
	"os"

	"github.com/charmbracelet/log"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type Config struct {
	GravityAcceleration float64 `json:"gravity_acceleration"`
	PendulumLength      float64 `json:"pendulum_length"`
	PendulumMass        float64 `json:"pendulum_mass"`
	SpringStiffness     float64 `json:"spring_stiffness"`
	DampingCoefficient  float64 `json:"damping_coefficient"`
	DistanceToSpring    float64 `json:"distance_to_spring"`
	InitialAngle1       float64 `json:"initial_angle1"`
	InitialAngle2       float64 `json:"initial_angle2"`
	MaxTime             float64 `json:"max_time"`
}

func calculate(config Config) (t, phi1, phi2, v1, v2 []float64) {
	omega1 := math.Sqrt(config.GravityAcceleration / config.PendulumLength)
	omega2 := math.Sqrt(config.GravityAcceleration/config.PendulumLength + 2*config.SpringStiffness*config.DistanceToSpring*config.DistanceToSpring/(config.PendulumMass*config.PendulumLength*config.PendulumLength))

	phi10 := (config.InitialAngle1 + config.InitialAngle2) / 2
	phi20 := (config.InitialAngle1 - config.InitialAngle2) / 2

	dt := 0.001

	t = make([]float64, 0, int(config.MaxTime/dt)+1)
	phi1 = make([]float64, 0, int(config.MaxTime/dt)+1)
	phi2 = make([]float64, 0, int(config.MaxTime/dt)+1)
	v1 = make([]float64, 0, int(config.MaxTime/dt)+1)
	v2 = make([]float64, 0, int(config.MaxTime/dt)+1)

	for i := 0.0; i <= config.MaxTime/dt; i++ {
		time := i * dt

		phi1t := (phi10*math.Cos(omega1*time) + phi20*math.Cos(omega2*time)) * math.Exp(-config.DampingCoefficient*time)
		phi2t := (phi10*math.Cos(omega1*time) - phi20*math.Cos(omega2*time)) * math.Exp(-config.DampingCoefficient*time)
		v1t := (-phi10*omega1*math.Sin(omega1*time) - phi20*omega2*math.Sin(omega2*time)) * math.Exp(-config.DampingCoefficient*time)
		v2t := (-phi10*omega1*math.Sin(omega1*time) + phi20*omega2*math.Sin(omega2*time)) * math.Exp(-config.DampingCoefficient*time)

		t = append(t, time)
		phi1 = append(phi1, phi1t)
		phi2 = append(phi2, phi2t)
		v1 = append(v1, v1t)
		v2 = append(v2, v2t)
	}

	fmt.Printf("Normal frequency omega_1: %0.4f Hz\n", omega1)
	fmt.Printf("Normal frequency omega_2: %0.4f Hz\n", omega2)

	return t, phi1, phi2, v1, v2
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

func main() {
	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal("failed to open config file", "err", err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("failed to unmarshal config file", "err", err)
	}

	t, phi1, phi2, v1, v2 := calculate(config)

	p := plot.New()

	p.Title.Text = "Coupled Pendulums"
	p.X.Label.Text = "Time (sec)"
	p.Y.Label.Text = "Angle (rad)"

	width := vg.Points(0.7)

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
	line2.Dashes = []vg.Length{vg.Points(10), vg.Points(10)}
	line2.Color = color.RGBA{B: 255, A: 255}

	p.Add(line1, line2)
	p.Legend.Add("First pendulum", line1)
	p.Legend.Add("Second pendulum", line2)

	p.X.Max = config.MaxTime / 2
	p.Y.Max = (sliceMax(phi1)+sliceMax(phi2))/2 + (sliceMin(phi1)+sliceMin(phi2))/2

	if err := p.Save(5*vg.Inch, 3*vg.Inch, "pendulum_angles_vs_time.png"); err != nil {
		log.Fatal(err)
	}

	p = plot.New()

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

	line3, err := plotter.NewLine(xy3)
	if err != nil {
		log.Fatal(err)
	}
	line3.Width = width
	line3.Color = color.RGBA{R: 255, A: 255}

	line4, err := plotter.NewLine(xy4)
	if err != nil {
		log.Fatal(err)
	}
	line4.Width = width
	line4.Dashes = []vg.Length{vg.Points(10), vg.Points(10)}
	line4.Color = color.RGBA{B: 255, A: 255}

	p.Add(line3, line4)
	p.Legend.Add("First pendulum", line3)
	p.Legend.Add("Second pendulum", line4)

	p.X.Max = config.MaxTime / 2
	p.Y.Max = (sliceMax(v1)+sliceMax(v2))/2 + (sliceMin(v1)+sliceMin(v2))/2

	if err := p.Save(5*vg.Inch, 3*vg.Inch, "pendulum_velocities_vs_time.png"); err != nil {
		log.Fatal(err)
	}
}
