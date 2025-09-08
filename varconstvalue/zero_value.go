package varconstvalue

import "fmt"

func ZeroValue() {
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)
}
