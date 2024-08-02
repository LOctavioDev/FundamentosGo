package main

import "fmt"

func main() {
	// ? OPERADORES DE COMPARACION
	// >, <, ==, !=, >=, <=
	fmt.Println(5 > 10)
	fmt.Println((5 * 5) == 10)
	fmt.Println(10 < 100)

	// ? OPERADORES LOGICOS
	// &&, ||
	var age uint = 85
	fmt.Println("ES ADULTO:", age >= 18 && age <= 60)
	fmt.Println("ES NINIO O ADULTO MAYOR:", age < 18 || age > 60)

	// ? OPERADOR UNARIO
	// '!'
	fmt.Println(!(4 == 4))

}
