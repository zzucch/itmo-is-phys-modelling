package config

import (
	"encoding/json"
	"errors"
)

type PendulumsConfig struct {
	GravityAcceleration float64 `json:"gravity_acceleration"`
	PendulumLength      float64 `json:"pendulum_length"`
	PendulumMass        float64 `json:"pendulum_mass"`
	SpringStiffness     float64 `json:"spring_stiffness"`
	DampingCoefficient  float64 `json:"damping_coefficient"`
	DistanceToSpring    float64 `json:"distance_to_spring"`
	InitialAngle1       float64 `json:"initial_angle1"`
	InitialAngle2       float64 `json:"initial_angle2"`
	MaxTime             float64 `json:"max_time"`
	TimeStep            float64 `json:"time_step"`
}

func ParsePendulumsData(data []byte) (*PendulumsConfig, error) {
	var cfg PendulumsConfig

	err := json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	if cfg.GravityAcceleration <= 0 ||
		cfg.PendulumLength <= 0 ||
		cfg.PendulumMass <= 0 ||
		cfg.SpringStiffness < 0 ||
		cfg.DampingCoefficient < 0 ||
		cfg.DistanceToSpring < 0 ||
		cfg.MaxTime <= 0 ||
		cfg.TimeStep <= 0 {
		return &cfg, errors.New("invalid configuration")
	}

	return &cfg, nil
}

