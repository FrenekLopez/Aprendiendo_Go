package mypackage

import (
	"fmt"
	"math/rand"
)

type User struct {
	Id   int
	Name string
}
type Employee struct {
	User
	Active bool
}
type Cashier interface {
	CalcTotal(item ...float32) float32
	deactive()
}
type Admin interface {
	DeactiveEmployee(c Cashier)
}

func NewEmployee(name string) *Employee {
	return &Employee{
		User: User{
			Id:   rand.Intn(1000),
			Name: name,
		},

		Active: true,
		fmt.Println(Id),
	}

}

func (e *Employee) CalcTotal(item ...float32) float32 {
	if !e.Active {
		fmt.Println("El cajero a sido desactivado")
		return 0
	}
	var sum float32
	for _, i := range item {
		sum += i
	}
	return sum * 1.15
}
func (e *Employee) deactive() {
	e.Active = false
}
func (e *Employee) DeactiveEmployee(c Cashier) {
	c.deactive()
}
