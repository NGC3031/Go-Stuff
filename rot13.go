package main

import (
	"fmt"
)

func main() {

	//message := "uv vagreangvbany fcnpr fgngvba"
	message:= "L fdph, L vdz, L frqtxhuhg"
	//	var c = 'd'
	for i := 0; i < len(message); i++ {
		c := message[i]
		//	fmt.Println(c)
		if c >= 'a' && c <= 'z' || c>='A' && c <='Z' {
			c = c -3
			if c > 'z' {
				c = c - 26
			}
		}

		fmt.Printf("%c", c)
	}
	fmt.Println("\n")
}
