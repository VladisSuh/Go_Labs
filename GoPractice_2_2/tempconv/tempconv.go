package tempconv

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Meter float64
type Feet float64
type Kilogram float64
type Pound float64

const (
	AbsoluteZeroC Celsius = -273.15
	Freeezing     Celsius = 0
	Boiling       Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%gC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gF", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
func (m Meter) String() string      { return fmt.Sprintf("%gM", m) }
func (fe Feet) String() string      { return fmt.Sprintf("%gF", fe) }
func (kg Kilogram) String() string  { return fmt.Sprintf("%gKg", kg) }
func (p Pound) String() string      { return fmt.Sprintf("%gP", p) }
