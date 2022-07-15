package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var piggyBank = 0
	var r int = 2

	for i := 0.0; i < 200000; i++ {
		rand.Seed(time.Now().UnixNano())
		r = rand.Intn(2) + 1
		//fmt.Printf("RND %v", r)
		switch r {
		case 1:
			piggyBank += 5
			fmt.Println("5c added")
		case 2:
			piggyBank += 10
			fmt.Println("10c added")
		case 3:
			piggyBank += 25
			fmt.Println("25c added")
		}
		//fmt.Printf(r)
		fmt.Printf("%vc\n", piggyBank)
		//fmt.Println(i)
		//fmt.Println(math.Abs(piggyBank - 20.0))
		//fmt.Println(math.Abs(piggyBank-20.0) < 0.05)
		if piggyBank >= 2000 {
			break
		}
	}
	var total float64 = float64(piggyBank) / 100
	fmt.Printf("\nAccount: $%.2f \n", total)
}
