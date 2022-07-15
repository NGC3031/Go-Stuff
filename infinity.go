// Note the random number gen is deterministic
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var degrees = 0
	var r = 0
	rand.Seed(time.Now().UnixNano())
	for {
		fmt.Println(degrees)
		degrees++
		if degrees >= 360 {
			degrees = 0
		}
		r = rand.Intn(360)
		fmt.Printf("Rand: %v \n", r)
		if r == degrees {
			fmt.Printf("Matched: %v \n", r)
			break
		}

	}

}
