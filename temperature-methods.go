package main

import "fmt"

type celsius float64
type kelvin float64
type farenheit float64

const kelv =273.15

// Kelvin to Celsius
func (k kelvin) celsius() celsius {
	return celsius(k - kelv)
}

// Celsius to Kelvin
func (c celsius) kelvin() kelvin {
	return kelvin(c + kelv)
}
// Farenheit to Celsius
func (f farenheit) farenheit() celsius {
	return celsius((f-32.0)*5.0/9.0)
}

func main() {
	var k kelvin = 294.0
	var c celsius = 127.0
	var f farenheit =78.0


	c = k.celsius()
	fmt.Print(k, "° K is ", c, "°C\n")

	k = c.kelvin()
	fmt.Print(c, "°C is ", k, "°K\n")

	c=f.farenheit()
	fmt.Print(f, "°F is ", c, "°C\n")

}
