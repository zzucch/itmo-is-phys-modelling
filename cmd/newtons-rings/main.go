package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/zzucch/itmo-is-phys-modelling/internal/calculation"
	"github.com/zzucch/itmo-is-phys-modelling/internal/config"
	"github.com/zzucch/itmo-is-phys-modelling/internal/plotting"
)

func main() {
	data, err := os.ReadFile("newtons_rings_config.json")
	if err != nil {
		log.Fatal("failed to open config file", "err", err)
	}

	cfg, err := config.ParseNewtonsRingsData(data)
	if err != nil {
		log.Fatal("failed to parse config file", "err", err)
	}

	result := calculation.CalculateNewtonsRings(*cfg)

	err = plotting.PlotIntensityDistribution(
		result,
		"rings_intensity_distribution.png")
	if err != nil {
		log.Fatal("failed to plot intensity distribution", "err", err)
	}
}
