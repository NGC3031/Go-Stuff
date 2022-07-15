package main

import "fmt"

func main() {

	year := 2018
	fmt.Printf("Type %T for %v\n", year, year)

	days := 365.2425
	fmt.Printf("Type %T for %[1]v\n", days)

	a := "text"
	fmt.Printf("Type %T for %[1]v\n", a)

	b := 42
	fmt.Printf("Type %T for %[1]v\n", b)

	c := 3.14
	fmt.Printf("Type %T for %[1]v\n", c)

	d := true
	fmt.Printf("Type %T for %[1]v\n", d)

	var green uint8 = 3
	fmt.Printf("%08b\n", green)
	green++
	fmt.Printf("%08b\n", green)

	//wrap aroun on integers
	var red uint8 = 255
	red = red + 67
	fmt.Println(red)

	var number int8 = 127
	number++
	fmt.Println(number)

}
