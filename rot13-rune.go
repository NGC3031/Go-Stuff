package main

import (
	"fmt"
"unicode/utf8"
)

func main() {

	//message := "uv vagreangvbany fcnpr fgngvba"
	//message:= "L fdph, L vdz, L frqtxhuhg"
        message:="¿Cómo estás?"
	for _, c:=range message {
                a, runeCount:=utf8.DecodeRuneInString(message)
                c=c-rune(runeCount)
		fmt.Printf("%c %c %v \n",c,a,runeCount)
		//	fmt.Println(c)
	/*if c >= 'a' && c <= 'z' || c>='A' && c <='Z' {
			c = c -3
			if c > 'z' {
				c = c - 26
			}
		}*/

		//fmt.Printf("\n - %c", c)
	}
	fmt.Println("\n")
}
