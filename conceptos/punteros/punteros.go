package main

import "fmt"

func main() {
	// ? EN CORTAS PALABRAS... LOS PUNTEROS SON VARIABLES QUE ALMACENAN LA DIRECCION DE MEMORIA DE UN VALOR

	var color string = "🔴"
	var pointerColor *string
	pointerColor = &color

	// PARA OBTENER LA DIRECCION DE MEMORIA USAMOS EL OPERADOR '&' ANTES DEL NOMBRE DE NUESRTRA VARIABLE
	fmt.Printf("TIPO: %T, VALOR: %s, DIRECCION: %v \n", color, color, &color)

	// PARA VER EL VALOR DE UN PUNTERO SE LE PONE UN '*' ANTES DEL PUNTERO... A ESTO SE LE LLAMA 'DESREFERENCIACION'
	fmt.Printf("TIPOS: %T, VALOR: %v, DESREFERENCIACION: %s \n", pointerColor, pointerColor, *pointerColor)

	//PARA ASIGNAR UN NUEVO VALOR HACEMOS LO MISMO QUE EN LA DESREFERENCIAICON
	*pointerColor = "🔵"

	fmt.Printf("TIPOS: %T, VALOR: %v, DESREFERENCIACION: %s \n", pointerColor, pointerColor, *pointerColor)
	
	// Y EL VALOR ORIGINAL DE LA VARIABLE 'color' TAMBIEN SE MODIFICA
	fmt.Printf("TIPO: %T, VALOR: %s, DIRECCION: %v \n", color, color, &color)


}
