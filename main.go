package main

import (
	"fmt"

	"github.com/freneklopez/aprendiendo_go/funciones"
)

// "github.com/freneklopez/aprendiendo_go/areas"
// "github.com/freneklopez/aprendiendo_go/areas"
// "github.com/freneklopez/aprendiendo_go/operadores"
// "github.com/freneklopez/aprendiendo_go/hello"
// "github.com/freneklopez/aprendiendo_go/operadores"
// "github.com/freneklopez/aprendiendo_go/complejos"
// "github.com/freneklopez/aprendiendo_go/print"

func main() {
	//hello.HolaMundo()
	//varConstValue.Constantes()
	//varConstValue.Variables()
	//varConstValue.ZeroValue()
	//operadores.Operadores()
	//areas.Rectangulos()
	//areas.Rectangulos()
	//areas.Circulo()
	//complejos.Complex()
	//print.Paquete()
	funciones.NormalFuncion("Erik hola")

	funciones.TripleArgumen(1, 2, "Texto")

	value := funciones.ReturnVAlue(5)
	fmt.Println("value", value)

	value1, _ := funciones.DoubleReturn(5)
	fmt.Println("Value1 ", value1)

}
