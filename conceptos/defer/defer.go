package main

import (
	"fmt"
	"os"
)

func main() {

	// ? DEFER ES UNA PALABRA RESERVADA EN GO QUE SIRVE PARA POSPONER LA EJECUCION DE DE UNA FUNCION HASTA QUE LA FUNCION QUE LA CONTIENE TERMINE.

	// ! SI HAY MAS DE UN DIFER, SE VAN A EJECUTAR EN EL ORDEN DE LIFO (Last In First Out) O EN OTRAS PALABRAS, EL ORDEN INVERSO DE COMO VAN DECLARADAS

	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	file, err := os.Create("index.js")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("console.log('Golang en mi father')"))
	if err != nil {
		fmt.Println(err)
		return
	}

}
