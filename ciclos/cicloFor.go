package ciclos

import (
	"fmt"
	"log"
	"strconv"
)

// For es un codigo que repite el proceso asta que se cumple la condicion.
func CicloFor() {
	//Convertir texto a numero
	value, err := strconv.Atoi("50")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value", value)

}
