package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/zzucch/itmo-is-phys-modelling/internal/calculation"
	"github.com/zzucch/itmo-is-phys-modelling/internal/config"
	"github.com/zzucch/itmo-is-phys-modelling/internal/plotting"
)

func main() {
	data, err := os.ReadFile("pendulums_config.json")
	if err != nil {
		log.Fatal("failed to open config file", "err", err)
	}

	cfg, err := config.Parse(data)
	if err != nil {
		log.Fatal("failed to parse config file", "err", err)
	}

	result := calculation.Calculate(*cfg)

	plotting.PlotAngles(
		result.T,
		result.Phi1,
		result.Phi2,
		cfg.MaxTime,
		"pendulums_angles_vs_time.png")

	plotting.PlotVelocities(
		result.T,
		result.V1,
		result.V2,
		cfg.MaxTime,
		"pendulums_velocities_vs_time.png")
}
