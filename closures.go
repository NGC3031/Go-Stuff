// From Get Programming with Go 14.6
// added my own comments adjustments etc
package main

import "fmt"

// kelvin declared as float64
type kelvin float64

// sensor function type
type sensor func() kelvin


func realSensor() kelvin {
	return 4 //added return 4 to see the flow through
	//In theory we could add whatever below
	//To-do: implement
}

// a real sensor
func calibrate(s sensor, offset kelvin) sensor {
	// Anonymous function with closure
	return func() kelvin {
		return s() + offset
	}
}

//Declares and returns an anonymous function
func main() {
	// sensor is of type func() kelvin, realSensor is a function
	// so a function is passed as a parameter along with offset
	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor())
}


