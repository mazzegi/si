package si

type Temperature float64

func Celsius(v float64) Temperature {
	return Temperature(v)
}

func Fahrenheit(v float64) Temperature {
	return Celsius((v - float64(32)) * 5.0 / 9.0)
}

func Kelvin(v float64) Temperature {
	return Celsius(v - 273.15)
}

func (t Temperature) Celsius() float64 {
	return float64(t)
}

func (t Temperature) Fahrenheit() float64 {
	return float64(t)*9.0/5.0 + 32.0
}

func (t Temperature) Kelvin() float64 {
	return float64(t) + 273.15
}
