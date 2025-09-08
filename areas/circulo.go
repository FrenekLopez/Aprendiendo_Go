package areas

import (
	"fmt"
	"math"
)

func Circulo() {
	radio := 3.0
	area := math.Pi * math.Pow(radio, 2)
	fmt.Println(area)
}
