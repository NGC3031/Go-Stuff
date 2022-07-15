package main

import (
	"fmt"
)

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	//cipherText := "ABCDEFG"
	keyword := "GOLANG"
	j := 0
	for i := 0; i < len(cipherText); i++ {
		//fmt.Printf("%c\t %v\t", cipherText[i], cipherText[i]-64)
		//fmt.Printf("%c\t %v\n", keyword[j], keyword[j]-64)
		c := cipherText[i]
		k := keyword[j] 
		r := (c-k+26)%26 + 'A'
		fmt.Printf("%c\t%v\n", r, r)
		j++
		if j >= len(keyword) {
			j = 0
		}
	}
}
