package config

import (
	"encoding/json"
	"errors"
)

type NewtonsRingsConfig struct {
	LensRadius     float64 `json:"lens_radius"`
	Wavelength     float64 `json:"wavelength"`
	SpectrumCenter float64 `json:"spectrum_center,omitempty"`
	SpectrumWidth  float64 `json:"spectrum_width,omitempty"`
}

func ParseNewtonsRingsData(data []byte) (*NewtonsRingsConfig, error) {
	var cfg NewtonsRingsConfig

	err := json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	if cfg.LensRadius <= 0 ||
		cfg.Wavelength < 0 ||
		cfg.SpectrumCenter < 0 {
		return &cfg, errors.New("invalid configuration")
	}

	return &cfg, nil
}
