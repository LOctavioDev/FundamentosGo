package main

import "fmt"

// ? A DIFERENCIA DE LOS LENUAJES TRADICIONALES COMO PYTHON, JAVASCRIPT, JAVA ETC... EN GO NO EXISTE COMO TAL EL CONCEPTO DE POO(PROGRAMACION ORIENTADA A OBJETOS), SI NO TIENE UNA FILOSOFIA DEL LENGUAJE MAS MODULAR, SE ENTIENDE COMO " structs "

// PARA HACERLO NECESITAMOS DECLARAR EL TIPO DE LA ESTRUCTURA, DESPUES EL NOMBRE Y AL FINAL LA PALABRA RESERVADA DEL LENGUAJE " structs "
// ESTO SERIA COMO EL EQUIVALENTE A UNA "CLASE" TRADICIONAL

// ! GENRALMENTE NO SE DECLARAN LAS STRUCTS EN LOS FUNCIONES, SI NO A NIVEL DE PAQUETE, PERO IGUAL LO PUEDES HACER DENTRO DE LA FUNCION

type Person struct {
	Name        string
	Age         uint8
	HasChildren bool
} // * ASI SE DECLARA UNA ESTRUCTURA

func main() {
	// ? IGUAL SE PUEDES HACER INSTANCIAS

	haziel := Person{
		Name:        "Eli Haziel",
		Age:         20,
		HasChildren: false,
	}

	// ! CUANDO CREAMOS UNA INSTANCIA, NO SE LE DICE " objeto ", EN DADO CASO DECIMOS QUE ES UNA INSTANCIA DE PERSONA O " un valor de tipo Persona "
	// ! NO SE USA LA TERMINOLOGIA DE " objeto " PORQUE GO NO ES UN LENGUAJE ORIENTADO A OBJETOS EN EL SENTIDO TRADICINAL

	// * TAMBIEN SE PUEDE DECLARA DE FORMA IMPLICITA
	mayra := Person{"Mayra", 30, true}

	// * TAMBIEN PUEDES CREAR NA INSTANCIA CON LOS UNICOS A TRIBUTOS QUE TU NECESITES
	antonio := Person{Age: 24}

	fmt.Printf("%+v", haziel)
	fmt.Printf("\n%+v", mayra)
	fmt.Printf("\n%+v\n", antonio)

	fmt.Println("==================================================")

	// DE LA MANERA TRADICIONAL PODEMOS ACCEDER A LAS PROPIEDADES DE LAS ESTRUCTURAS
	fmt.Printf("%+v\n", haziel.Name)
	fmt.Printf("%+v\n", mayra.HasChildren)
	fmt.Printf("%+v\n", antonio.Age)

	fmt.Println("==================================================")

	// ? PARA USAR LOS PUNTEROS CON LAS ESTRCUTURAS SE HACE DE IGUAL MANERA CON EL " & "
	yayo := &haziel

	// ? PARA MODFICAR DESDE EL PUNTERO LO PUEDES ENCERRAMOS EN UNOS PARENTESIS (*puntero) O HACERLO DE LA MANERA TRADICIONAL
	yayo.Age = 21
	fmt.Printf("%+v\n", haziel)
	fmt.Printf("%+v", yayo)

}
