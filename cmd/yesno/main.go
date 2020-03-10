package main

import (
	"fmt"
	"github.com/xyproto/dialog"
)

func main() {
	d := dialog.New(80, 24)
	answeredYes, err := d.YesNo("Do you want cake?")
	if err != nil {
		panic(err)
	}
	if answeredYes {
		fmt.Println("You answered: yes")
		fmt.Println("The cake is a lie. Haha!")
	} else {
		fmt.Println("You answered: no")
		fmt.Println("Fine.")
	}
}
