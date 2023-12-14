package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		epsX                   = (xmax - xmin) / width
		epsY                   = (ymax - ymin) / height
	)
	offX := []float64{-epsX, epsX}
	offY := []float64{-epsY, epsY}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			subPixels := make([]color.Color, 0)
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					z := complex(x+offX[i], y+offY[j])
					subPixels = append(subPixels, mandelbrot(z))
				}
			}
			img.Set(px, py, avg(subPixels))
		}
	}
	f, _ := os.Create("output.png")
	png.Encode(f, img)
}

var palettes = [...]color.RGBA{
	{66, 30, 15, 255},
	{25, 7, 26, 255},
	{9, 1, 47, 255},
	{4, 4, 73, 255},
	{0, 7, 100, 255},
	{12, 44, 138, 255},
	{24, 82, 177, 255},
	{57, 125, 209, 255},
	{134, 181, 229, 255},
	{211, 236, 248, 255},
	{241, 233, 191, 255},
	{248, 201, 95, 255},
	{255, 170, 0, 255},
	{204, 128, 0, 255},
	{153, 87, 0, 255},
	{106, 52, 3, 255},
}

func avg(colors []color.Color) color.Color {
	var r, g, b, a uint16
	n := len(colors)
	for _, c := range colors {
		r_, g_, b_, a_ := c.RGBA()
		r += uint16(r_ / uint32(n))
		g += uint16(g_ / uint32(n))
		b += uint16(b_ / uint32(n))
		a += uint16(a_ / uint32(n))
	}
	return color.RGBA64{r, g, b, a}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			switch {
			case n > 50:
				return color.RGBA{66, 30, 15, 255}
			default:
				return palettes[n%16]
			}
		}
	}
	return color.Black
}
