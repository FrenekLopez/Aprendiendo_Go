package print

import (
	"fmt"
)

func Paquete() {
	//Println (Sirve para el salto de linea)
	fmt.Println("Hello word")
	fmt.Println("Hello word")

	//Printf (Sirve para agregar dos variables en una misma linea)
	nombre := "Eric"
	edad := 30
	fmt.Printf("%s es mi nombre y tengo %d \n", nombre, edad)
	fmt.Printf("%v es mi nombre y tengo %v \n", nombre, edad)

	// Srintf(Sirve para agregar esa linea de comando a una variable e imprimirla)
	menssege := fmt.Sprintf("%s es mi nombre y tengo %d", nombre, edad)
	fmt.Println(menssege)

	// Para imprimir el tipo de dato al que pertenece

	fmt.Printf("la variable %v es un %T\n", nombre, nombre)

}
