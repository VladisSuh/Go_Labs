package main

import (
	"GoPractice_2_2/tempconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %c\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		m := tempconv.Meter(t)
		fe := tempconv.Feet(t)
		kg := tempconv.Kilogram(t)
		p := tempconv.Pound(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToK(f), k, tempconv.KToF(k))
		fmt.Printf("%s = %s, %s = %.1f\n",
			c, tempconv.CToK(c), k, tempconv.KToC(k))
		fmt.Printf("%s = %s, %s = %s\n",
			m, tempconv.MToFe(m), fe, tempconv.FeToM(fe))
		fmt.Printf("%s = %s, %s = %s\n",
			kg, tempconv.KgToP(kg), p, tempconv.PToKg(p))
	}
}
