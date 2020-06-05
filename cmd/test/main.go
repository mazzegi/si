package main

import (
	"fmt"

	"github.com/mazzegi/si"
)

func main() {
	l := si.Metre(1)
	fmt.Println("1 * si.Metre")
	fmt.Println("in m: ", l.Metre())
	fmt.Println("in inch: ", l.Inch())
	fmt.Println("in km: ", l.KiloMetre())

	l = si.Inch(3)
	fmt.Println("\n3 * si.Inch")
	fmt.Println("in m: ", l.Metre())
	fmt.Println("in inch: ", l.Inch())
	fmt.Println("in km: ", l.KiloMetre())

	fmt.Println("\ntemps")
	t := si.Celsius(25)
	fmt.Println("°C: ", t.Celsius())
	fmt.Println("K: ", t.Kelvin())
	fmt.Println("°F: ", t.Fahrenheit())
}
