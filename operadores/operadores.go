package operadores

import "fmt"

var x int = 5
var y int = 5

func Operadores() {

	//Suma
	suma := x + y
	fmt.Println("Suma", suma)
	//Resta
	resta := x + y
	fmt.Println("Resta", resta)
	//Divivion
	divicion := x / y
	fmt.Println("Divicion", divicion)
	//Multiplicacion
	mult := x * y
	fmt.Println("Multiplicacion", mult)
	//Incremental
	x++
	fmt.Println("Incremental", x)
	//Decremental
	x--
	fmt.Println("Decremental", x)
}
