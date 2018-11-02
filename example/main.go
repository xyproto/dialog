package main

import (
	"fmt"

	"github.com/xyproto/dw"
)

func main() {
	d := dw.NewDialog(80, 20)
	d.MsgBox("hi")
	answer := d.Menu("hi", 6, map[string]string{"a": "A", "b": "B"})
	fmt.Printf("\n\n\n\n%s\n", answer)
}
