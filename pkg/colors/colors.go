package colors

import (
	"image/color"

	"github.com/converged-computing/distributed-fractal/pkg/algorithm"
)

// interpolate a color palette code into a list of RGBA
// this is taken from (MIT LICENSE):
// https://github.com/esimov/gobrot/blob/master/internal/brot/main.go#L27
func InterpolateColors(paletteCode *string, numberOfColors int) []color.RGBA {
	var factor float64
	var steps []float64
	var cols []uint32
	var interpolated []uint32
	var interpolatedColors []color.RGBA

	for _, v := range ColorPalettes {
		factor = 1.0 / float64(numberOfColors)
		if v.Keyword != *paletteCode || paletteCode == nil {
			continue
		}
		for index, col := range v.Colors {
			if col.Step == 0.0 && index != 0 {
				stepRatio := float64(index+1) / float64(len(v.Colors))
				step := float64(int(stepRatio*100)) / 100 // truncate to 2 decimal precision
				steps = append(steps, step)
			} else {
				steps = append(steps, col.Step)
			}
			r, g, b, a := col.Color.RGBA()
			r /= 0xff
			g /= 0xff
			b /= 0xff
			a /= 0xff
			uintColor := r<<24 | g<<16 | b<<8 | a
			cols = append(cols, uintColor)
		}

		var min, max, minColor, maxColor float64
		if !(len(v.Colors) == len(steps) && len(v.Colors) == len(cols)) {
			continue
		}
		for i := 0.0; i <= 1; i += factor {
			for j := 0; j < len(v.Colors)-1; j++ {
				if !(i >= steps[j] && i < steps[j+1]) {
					continue
				}
				min = steps[j]
				max = steps[j+1]
				minColor = float64(cols[j])
				maxColor = float64(cols[j+1])
				uintColor := algorithm.CosineInterpolation(
					maxColor,
					minColor,
					(i-min)/(max-min),
				)
				interpolated = append(interpolated, uint32(uintColor))
			}
		}

		for _, pixelValue := range interpolated {
			r := pixelValue >> 24 & 0xff
			g := pixelValue >> 16 & 0xff
			b := pixelValue >> 8 & 0xff
			a := 0xff

			interpolatedColors = append(
				interpolatedColors,
				color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)},
			)
		}
	}

	return interpolatedColors
}
