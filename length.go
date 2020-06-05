package si

type Length float64

const (
	metre     float64 = 1
	inch      float64 = 0.0254
	kiloMetre float64 = 1000
)

func Metre(v float64) Length {
	return Length(v)
}

func KiloMetre(v float64) Length {
	return Metre(v * kiloMetre)
}

func Inch(v float64) Length {
	return Metre(v * inch)
}

func (l Length) Metre() float64 {
	return float64(l)
}

func (l Length) KiloMetre() float64 {
	return float64(l) / kiloMetre
}

func (l Length) Inch() float64 {
	return float64(l) / inch
}
