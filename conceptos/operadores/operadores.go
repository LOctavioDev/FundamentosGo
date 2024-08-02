package main

import "fmt"

// LOS OPERADORES ARITMETICOS Y DE ASIGNACION SON COMO EN CASI TODOS LOS LENGAUJES DE PROGRAMACION

func main() {
	// ? OPERADORES ARITMETICOS
	// ESTA ES LA JERARQUIA DE OPERACIONES: (), *, /, %, +, -
	var a = 20 + 3*2/(5+6)
	fmt.Println(a)

	// ? OPERADORES DE ASIGANCION
	// =, +=, -=, *=, /=, %=
	var b int = 5
	b += 5
	b -= 5
	b *= 1
	b /= 5
	b %= 1
	fmt.Println(b)

	// ? GO TAMBIEN TIENE LA DECLARACION POST-INCREMMENTO Y POST-DECREMENTO: ++, --
	// ! NO SON UNA EXPRESION SINO UNA DECALRACION

	/*
		! NO SE PUEDE HACER ESTO:) ADIFERENCIA DE JAVASCRIPT POR EJMEPLO
			var c int = 11
			fmt.Println(c++)
	*/

	// * ASI SI SE PUEDE
	var c int = 11
	c++
	c--
	fmt.Println(c)
}
