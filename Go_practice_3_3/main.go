package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

type CornerType int

const (
	middle CornerType = 0
	peak   CornerType = 1
	valley CornerType = 2
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ct, err := corner(i+1, j)
			if err != nil {
				continue
			}
			bx, by, ct1, err := corner(i, j)
			if err != nil {
				continue
			}
			cx, cy, ct3, err := corner(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, ct2, err := corner(i+1, j+1)
			if err != nil {
				continue
			}

			var color string
			if ct == peak || ct1 == peak || ct2 == peak || ct3 == peak {
				color = "red"
			} else if ct == valley || ct1 == valley || ct2 == valley || ct3 == valley {
				color = "blue"
			} else {
				color = "grey"
			}

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: %s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, CornerType, error) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z, ct := f(x, y)

	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, 0, fmt.Errorf("invalid value")
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, ct, nil
}

func f(x, y float64) (float64, CornerType) {
	d := math.Hypot(x, y)
	ct := middle

	if math.Abs(d-math.Tan(d)) < 3 {
		ct = peak
		if 2*(math.Sin(d)-d*math.Cos(d))-d*d*math.Sin(d) > 0 {
			ct = valley
		}
	}
	return math.Sin(d) / d, ct
}
