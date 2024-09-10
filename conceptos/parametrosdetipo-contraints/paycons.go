package main

import "fmt"

type MyInt int
type MyIntV2 int

func main() {
	// ? LOS PARAMETROS DE TIPO NOS PERMITE AGREGAR RESTRICCIONES DE SIMPLE TIPOS DE DATOS... restricciones

	var num1 MyInt = 2
	var num2 MyInt = 6

	var num3 MyIntV2 = 10
	var num4 MyIntV2 = 20

	// fmt.Println(sumar[string]("hola", "mi amor")) // ! no se puede
	fmt.Println(sumar(1, 2, 3, 4, 5))
	fmt.Println(sumar(1, 2.2, 3.3))
	fmt.Println(sumar(num1, num2))
	fmt.Println(sumar(num3, num4))
}

// EL PROBLEMA DE ANY ES QUE NO SE PUEDE SUMAR +...
func sumar[T ~int | float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}

	return total
}
