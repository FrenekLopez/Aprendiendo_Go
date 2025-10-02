package condicionales

import "fmt"

func MySwich() {

	switch modulo := 160 % 2; modulo {
	case 0:
		fmt.Println("El resultado es un par ", modulo)
	default:
		fmt.Println("El numero es impar", modulo)
	}
	//switch sin condicion
	dia := "viernes"
	switch dia {
	case "martes":
		fmt.Println("Inicio de semana")
	case "viernes":
		fmt.Println("¡Fin de semana!")
	default:
		fmt.Println("Día normal")
	}
}
