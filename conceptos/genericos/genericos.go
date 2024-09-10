package main

import "fmt"

func main() {
	// ? LOS GENERICOS COMO TAL ES UN CONCEPTO QUE PERMITE ESCRIBIR FUNCIONES Y ESTRUCTURAS DE DATOS DE MANERA SEGURA Y SIN TENER QUE DEFINIR EL TIPO EXCATO DE ANTEMANO.

	// * EN ESTE EJEMPLO FUNCIONA BIEN, SIEMPRE PODEMOS PASAR VALORES COMO CADENAS...
	PrintList("Mi", "buen", "amoooor")

	// * SIN EMBARGO QUE PASA SI AHORA NO QUIERO SOLO MOSTRAR VALORES DE CADENA, SI NO TAMBIEN NUMEROS U OTRO TIPO DE DATO
	// ! PrintList(1, 2, 3, 4) cannot use 4 (untyped int constant) as string value in argument to PrintListcompilerIncompatibleAssign

	PrintList("Patada", "de", "canguro", "golpe", "duro")
	PrintList(1, 2, 3)

}

// ? GO NOS DA UN TIPO DE DATO "any" PARA MANEJAR EN ESOS CASO QUE RECIBA CUALQUIER TIPO DE DATO... ASI ES COMO EN TYPESCRIPT

func PrintList(list ...any) {
	for _, item := range list {
		fmt.Println(item)
	}
}
