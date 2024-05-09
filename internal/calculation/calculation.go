package calculation

import (
	"fmt"
	"math"

	"github.com/zzucch/itmo-is-phys-modelling/internal/config"
)

func Calculate(cfg config.Config) (t, phi1, phi2, v1, v2 []float64) {
	omega1, omega2 := calculateFrequencies(cfg)
	phi10, phi20 := calculateInitialAngles(cfg.InitialAngle1, cfg.InitialAngle2)

	dt := 0.001

	t = make([]float64, 0, int(cfg.MaxTime/dt)+1)
	phi1 = make([]float64, 0, int(cfg.MaxTime/dt)+1)
	phi2 = make([]float64, 0, int(cfg.MaxTime/dt)+1)
	v1 = make([]float64, 0, int(cfg.MaxTime/dt)+1)
	v2 = make([]float64, 0, int(cfg.MaxTime/dt)+1)

	for i := 0.0; i <= cfg.MaxTime/dt; i++ {
		time := i * dt

		phi1t, phi2t, v1t, v2t := calculateValuesAtTime(
			time, phi10, phi20, omega1, omega2, cfg.DampingCoefficient)

		t = append(t, time)
		phi1 = append(phi1, phi1t)
		phi2 = append(phi2, phi2t)
		v1 = append(v1, v1t)
		v2 = append(v2, v2t)
	}

	printResults(omega1, omega2)

	return t, phi1, phi2, v1, v2
}

func calculateFrequencies(cfg config.Config) (float64, float64) {
	omega1 := math.Sqrt(cfg.GravityAcceleration / cfg.PendulumLength)

	omega2 := math.Sqrt(cfg.GravityAcceleration/cfg.PendulumLength +
		2*cfg.SpringStiffness*cfg.DistanceToSpring*cfg.DistanceToSpring/
			(cfg.PendulumMass*cfg.PendulumLength*cfg.PendulumLength))

	return omega1, omega2
}

func calculateInitialAngles(angle1, angle2 float64) (float64, float64) {
	phi10 := (angle1 + angle2) / 2
	phi20 := (angle1 - angle2) / 2

	return phi10, phi20
}

func calculateValuesAtTime(
	time, phi10, phi20, omega1, omega2, dampingCoefficient float64,
) (float64, float64, float64, float64) {
	phi1t := (phi10*math.Cos(omega1*time) +
		phi20*math.Cos(omega2*time)) * math.Exp(-dampingCoefficient*time)

	phi2t := (phi10*math.Cos(omega1*time) -
		phi20*math.Cos(omega2*time)) * math.Exp(-dampingCoefficient*time)

	v1t := (-phi10*omega1*math.Sin(omega1*time) -
		phi20*omega2*math.Sin(omega2*time)) * math.Exp(-dampingCoefficient*time)

	v2t := (-phi10*omega1*math.Sin(omega1*time) +
		phi20*omega2*math.Sin(omega2*time)) * math.Exp(-dampingCoefficient*time)

	return phi1t, phi2t, v1t, v2t
}

func printResults(omega1, omega2 float64) {
	fmt.Println("Results:")
	fmt.Printf("Normal frequency omega_1: %0.4f Hz\n", omega1)
	fmt.Printf("Normal frequency omega_2: %0.4f Hz\n", omega2)
}
