package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 3)
	c <- "Email1"
	c <- "Email2"
	c <- "Email3"

	// Close
	close(c)

	//fmt.Println(len(c), cap(c))

	//c <- "Mensaje 3"
	for saludo := range c {
		fmt.Println(saludo)
	}
	//Select
	email1 := make(chan string)
	email2 := make(chan string)
	email3 := make(chan string)

	go message("Mensaje1", email1)
	go message("Mensaje2", email2)
	go message("Mensaje3", email3)
	for i := 0; i < 3; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		case m3 := <-email3:
			fmt.Println("Email recibido de email3", m3)
		}

	}
}
