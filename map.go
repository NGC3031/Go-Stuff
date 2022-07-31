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
	// Now try and read a key that doesnt exist should give 0
	moon := temperature["Moon"]
	fmt.Println(moon)
	// Comma ok syntax to test a key
	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("On average the moon is %vo C.\n", moon)
	} else {
		fmt.Println("Where is the moon?")
	}
	// we can delete a key
	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}
	planetsMarkII := planets

	planets["Earth"] = "whoops"
	fmt.Println(planets)
	fmt.Println(planetsMarkII)
	delete(planets, "Earth")
	fmt.Println(planetsMarkII)
}
