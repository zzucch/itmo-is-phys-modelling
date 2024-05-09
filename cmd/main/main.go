package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/zzucch/itmo-is-phys-modelling/internal/calculation"
	"github.com/zzucch/itmo-is-phys-modelling/internal/config"
	"github.com/zzucch/itmo-is-phys-modelling/internal/plotting"
)

func main() {
	data, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal("failed to open config file", "err", err)
	}

	cfg, err := config.Parse(data)
	if err != nil {
		log.Fatal("failed to parse config file", "err", err)
	}

	t, phi1, phi2, v1, v2 := calculation.Calculate(*cfg)

	plotting.PlotAngles(
		t, phi1, phi2, cfg.MaxTime, "pendulum_angles_vs_time.png")

	plotting.PlotVelocities(
		t, v1, v2, cfg.MaxTime, "pendulum_velocities_vs_time.png")
}
