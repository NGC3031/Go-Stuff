package main

import (
	"fmt"
)

func main() {
	yesNo := "yes"
	launch := false

	switch yesNo {
	case "yes":
		launch = (yesNo == "yes")
        
	default:
		launch = false
 	}

	fmt.Println("Ready for launch:", launch)
}
