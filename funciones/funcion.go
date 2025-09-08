package funciones

import "fmt"

func NormalFuncion(message string) {
	fmt.Println(message)

}

func TripleArgumen(a, b int, c string) {
	fmt.Println(a, b, c)

}

func ReturnVAlue(a int) int {
	return a * 2

}
func DoubleReturn(a int) (b, c int) {
	return a, a * 2
}
