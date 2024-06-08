package plotting

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/zzucch/itmo-is-phys-modelling/internal/calculation"
)

func PlotIntensityDistribution(
	result calculation.NewtonsRingsResult, filename string,
) error {
	img := image.NewRGBA(image.Rect(0, 0, result.Width, result.Height))

	for y := 0; y < result.Height; y++ {
		for x := 0; x < result.Width; x++ {
			intensity := result.Intensities[y][x]
			r := uint8(intensity[0] * 255)
			g := uint8(intensity[1] * 255)
			b := uint8(intensity[2] * 255)
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, img)
}
