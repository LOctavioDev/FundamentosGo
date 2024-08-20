package main

import "fmt"

func main() {

	// ? EN GO PARA DECLARA UN ARREGLO SE HACE DECLARANDO EL NOMBRE DE LA VARIABLE, SU TAMANIO Y SU TIPO DE DATOS.
	var flags [3]string
	flags[0] = "🏴"
	flags[1] = "🏳️"
	flags[2] = "🏁"

	// ! EN GO, LOS ARREGLOS SON DE TAMANIO ESTATICO
	// flags[3] = "x" ESTO NO ES POSIBLE

	fmt.Println(flags)

	// ? LOS ARRAYS LITERALES SON UNA MANERA MAS PRACTICA PARA MANEJAR LOS ARRYAS
	// * USANDO EL OPERADOR DE ASIGNACION CORTA

	// ! UNA MANERA DE QUE GO INFIERA EL TAMANIO DE MI ARREGLO ES CON EL OPERADOR DE " ... "
	fruits := [...]string{"🍊", "🍎", "🍌"}

	// DE TODAS MANERAS SIGUE SIENDO DE TAMANIO FIJO, ASI QUE NO PODEMOS HACER LO SIGUIENTE
	// fruits[5] = "d"  => NO SE PUEDE

	fmt.Println(fruits)
}
