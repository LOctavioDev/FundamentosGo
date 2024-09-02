package main

import "fmt"

// ? LAS FUNCIONES VARIATICAS SON AQUELLAS QUE NOS DEAJAN PASAR UN NUMERO N DE PARAMETROS
// ? AL IGUAL QUE EN JAVASCRIPT, LAS FUNCIONES ANONIMAS SIMPLEMENTE SON FUNCIONES QUE NO TIENEN NOMBRE

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5))

	// * SE PUEDEN GUARDAR EN VARIABLES
	greet := func ()  {
		fmt.Println("👋 console.log(test) ALV JavaScirpt")
	}
	greet()

	// * Y SI SOLO QUIERO QUE SE EJECUEN SIN ASIGNARLE A UNA VARIABLE SOLO PONEMOS LOS PARENTESIS AL FINAL
	// ! PUEDEN RECIBIR PARAMETROS O NO AMBOS TIPOS DE FUNCIONES

	func (language string)  {
		fmt.Printf("%v, es mejor babys", language)
	}("C++")

}

func sum(nums ...int) (total int) {
	for _, num := range nums {
		total += num
	}

	return
}
