package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		x := parse(r.Form["x"], 0)
		y := parse(r.Form["y"], 0)
		zoom := parse(r.Form["zoom"], 0)
		renderPNG(w, x, y, zoom)
	}
	http.HandleFunc("/", handler)

	fmt.Println("http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func renderPNG(out io.Writer, x, y, zoom float64) {
	const (
		width, height = 1024, 1024
	)

	m := math.Exp2(1 - zoom)
	xmin, xmax := x-m, x+m
	ymin, ymax := y-m, y+m

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img)
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

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palettes[n%16]
		}
	}
	return color.Black
}

func parse(array []string, defaultValue float64) float64 {
	if len(array) < 1 {
		return defaultValue
	}
	value, err := strconv.ParseFloat(array[0], 64)
	if err != nil {
		return defaultValue
	}
	return value
}
