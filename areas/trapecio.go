package areas

import "fmt"

func Trapecio() {
	var baseMenor float64 = 5
	var baseMayor float64 = 8
	var altura float64 = 5
	area := ((baseMenor + baseMayor) * altura) / 2
	fmt.Println(area)
}
