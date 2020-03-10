package main

import (
	"fmt"
	"github.com/xyproto/dialog"
)

func main() {
	d := dialog.New(80, 24)
	answer, err := d.Menu("Choose your number", 6, map[string]string{"2": "More than one", "7": "A quality choice", "13": "A controversial option"})
	if err != nil {
		panic(err)
	}
	fmt.Println("Your number: " + answer)
}
