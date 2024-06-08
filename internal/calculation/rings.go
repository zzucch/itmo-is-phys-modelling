package calculation

import (
	"math"

	"github.com/zzucch/itmo-is-phys-modelling/internal/config"
)

type NewtonsRingsResult struct {
	Intensities [][][3]float64
	Width       int
	Height      int
}

func CalculateNewtonsRings(cfg config.NewtonsRingsConfig) NewtonsRingsResult {
	lensRadius := cfg.LensRadius
	imageSize := 5000
	center := float64(imageSize) / 2.0

	intensities := make([][][3]float64, imageSize)
	for i := range intensities {
		intensities[i] = make([][3]float64, imageSize)
	}

	for y := 0; y < imageSize; y++ {
		for x := 0; x < imageSize; x++ {
			r := math.Sqrt(
				(float64(x)-center)*(float64(x)-center) +
					(float64(y)-center)*(float64(y)-center))

			r = r * lensRadius / center

			if cfg.SpectrumWidth == 0 {
				intensities[y][x] = calculateMonochromaticIntensity(
					r, cfg.Wavelength)
			} else {
				intensities[y][x] = calculateQuasiMonochromaticIntensity(
					r, cfg.SpectrumCenter, cfg.SpectrumWidth)
			}
		}
	}

	return NewtonsRingsResult{
		Intensities: intensities,
		Width:       imageSize,
		Height:      imageSize,
	}
}

func calculateMonochromaticIntensity(r, wavelength float64) [3]float64 {
	intensity := 0.5 * (1 + math.Cos(2*math.Pi*r/wavelength))
	return [3]float64{
		intensity,
		intensity,
		intensity,
	}
}

func calculateQuasiMonochromaticIntensity(r, center, width float64) [3]float64 {
	const samplesAmount = 10
	var totalIntensity [3]float64
	step := width / float64(samplesAmount)

	for i := 0; i < samplesAmount; i++ {
		wavelength := center - width/2 + step*float64(i)
		intensity := calculateMonochromaticIntensity(r, wavelength)
		color := WavelengthToRGB(wavelength)

		for j := 0; j < 3; j++ {
			totalIntensity[j] += intensity[j] * color[j]
		}
	}

	for j := 0; j < 3; j++ {
		totalIntensity[j] /= float64(samplesAmount)
	}

	return totalIntensity
}

func WavelengthToRGB(wavelength float64) [3]float64 {
	var R, G, B float64

	switch {
	case wavelength >= 380 && wavelength < 440:
		R = -(wavelength - 440) / (440 - 380)
		G = 0.0
		B = 1.0
	case wavelength >= 440 && wavelength < 490:
		R = 0.0
		G = (wavelength - 440) / (490 - 440)
		B = 1.0
	case wavelength >= 490 && wavelength < 510:
		R = 0.0
		G = 1.0
		B = -(wavelength - 510) / (510 - 490)
	case wavelength >= 510 && wavelength < 580:
		R = (wavelength - 510) / (580 - 510)
		G = 1.0
		B = 0.0
	case wavelength >= 580 && wavelength < 645:
		R = 1.0
		G = -(wavelength - 645) / (645 - 580)
		B = 0.0
	case wavelength >= 645 && wavelength <= 780:
		R = 1.0
		G = 0.0
		B = 0.0
	default:
		R = 0.0
		G = 0.0
		B = 0.0
	}

	return [3]float64{R, G, B}
}
