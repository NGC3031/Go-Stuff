// From get Programming with Go
package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15,
		// Composite literals are key-value pairs for maps
		"Mars": -65,
	}
	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %vo C.\n", temp)
	temperature["Earth"] = 16
	temperature["Venus"] = 464
	fmt.Println(temperature)
}
