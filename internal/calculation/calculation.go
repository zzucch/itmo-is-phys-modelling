package calculation

import (
	"fmt"
	"math"

	"github.com/zzucch/itmo-is-phys-modelling/internal/config"
)

type CalculationResult struct {
	T    []float64
	Phi1 []float64
	Phi2 []float64
	V1   []float64
	V2   []float64
}

func Calculate(cfg config.Config) CalculationResult {
	size := int(math.Ceil(cfg.MaxTime / cfg.TimeStep))
	result := CalculationResult{
		T:    make([]float64, 0, size),
		Phi1: make([]float64, 0, size),
		Phi2: make([]float64, 0, size),
		V1:   make([]float64, 0, size),
		V2:   make([]float64, 0, size),
	}

	omega1, omega2 := calculateFrequencies(cfg)
	printResults(omega1, omega2)

	phi1Initial, phi2Initial := calculateInitialAngles(cfg.InitialAngle1, cfg.InitialAngle2)

	for time := 0.0; time <= cfg.MaxTime; time += cfg.TimeStep {
		phi1t, phi2t, v1t, v2t := calculateValuesAtTime(
			time, phi1Initial, phi2Initial, omega1, omega2, cfg.DampingCoefficient)

		appendToResult(&result, time, phi1t, phi2t, v1t, v2t)
	}

	return result
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

func appendToResult(
	result *CalculationResult, time, phi1, phi2, v1, v2 float64,
) {
	result.T = append(result.T, time)
	result.Phi1 = append(result.Phi1, phi1)
	result.Phi2 = append(result.Phi2, phi2)
	result.V1 = append(result.V1, v1)
	result.V2 = append(result.V2, v2)
}

func printResults(omega1, omega2 float64) {
	fmt.Println("Results:")
	fmt.Printf("Normal frequency omega_1: %0.4f Hz\n", omega1)
	fmt.Printf("Normal frequency omega_2: %0.4f Hz\n", omega2)
}
