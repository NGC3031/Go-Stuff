// Demo of sorting on a slice
package main

import (
	"fmt"
	"sort"
)

// dump slice length, capacity, and contents
func dump(label string, slice []string) {
fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice),cap(slice), slice)
}


func main() {
        // Declare a slice
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
        planets =append(planets,"Krypton")
	//Sorts planets alphabetically
	sort.StringSlice(planets).Sort()
	fmt.Println(planets)
dump("Planets Slice", planets)

}
