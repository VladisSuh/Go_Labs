package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

func KToC(k Kelvin) Celsius { return Celsius(k) + AbsoluteZeroC }

func FToK(f Fahrenheit) Kelvin {
	FahrenheitToCelsius := Celsius((f - 32) * 5 / 9)
	return Kelvin(FahrenheitToCelsius - AbsoluteZeroC)
}

func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit((k+Kelvin(AbsoluteZeroC))*9/5 + 32)
}

func MToFe(m Meter) Feet {
	return Feet(m * 3.28084)
}

func FeToM(fe Feet) Meter {
	return Meter(fe * 0.3048)
}

func KgToP(kg Kilogram) Pound {
	return Pound(kg * 2.20462)
}

func PToKg(p Pound) Kilogram {
	return Kilogram(p * 0.453592)
}
