package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	piggyBank := 0.0
	var r int = 2

	for i := 0.0; i < 200000; i++ {
		rand.Seed(time.Now().UnixNano())
		r = rand.Intn(2) + 1
		//fmt.Printf("RND %v", r)
		switch r {
		case 1:
			piggyBank += 0.05
			fmt.Println("5c added")
		case 2:
			piggyBank += 0.1
			fmt.Println("10c added")
		case 3:
			piggyBank += 0.25
			fmt.Println("25c added")
		}
		//fmt.Printf(r)
		fmt.Printf("$ %.2f \n", piggyBank)
		//fmt.Println(i)
		//fmt.Println(math.Abs(piggyBank - 20.0))
		//fmt.Println(math.Abs(piggyBank-20.0) < 0.05)
		if math.Abs(piggyBank-20.0) <= 0.05 {
			break
		}
	}
	fmt.Printf("\nAccount: %.2f \n", piggyBank)
}
