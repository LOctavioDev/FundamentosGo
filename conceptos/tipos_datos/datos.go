package main

import "fmt" // RECORDEMOS QUE 'fmt' ES UN PAUQUETE ESTANDAR DE GO QUE PROPORCIONA FUNCIONES PARA FORMATEAR E/S (entrada / salida)

func main() {
	// EN GO TENEMOS LOS TRES PRINCIPALES TIPOS DE DATOS
	// 'string', 'bool', 'numeric'

	var a bool = true
	var b string = "Octavio"
	// LOS VERBOS DE FORMATO SON SECUENCIAS DE CARACTERES UTILIZADAS EN CADENAS DE FORMATO
	//INICIAN CON '%' Y UNA LETRA CLAVE. ESTAN EN LA DOCUEMNTACION OFICAL DE GO
	//* https://go.dev/play/p/TLwZJSrapB

	fmt.Printf("TIPO: %T, VALOR: %v \n", a, a)
	fmt.Printf("TIPO: %T, VALOR: %v \n", b, b)

	// ? LOS TIPOS DE DATOS NUMERIC SE DIVEN EN MAS SUBTIPOS. HAY UN ARCHIVO .MD DONDE LOS PUEDES VISUALIZAR MAS A DETALLE

	// ACA UN EJEMPLO DE BUEN USO
	// NUNCA HABRA UNA EDAD DE UNA PERSONA QUE SEA NEGTIVA O SUPERE LOS 255 ANIOS
	var age uint8 = 20
	fmt.Printf("TIPO: %T, VALOR: %v \n", age, age)

	//TAMBIEN TENEMOS 'alias' POR EJEMPLO 'byte' QUE ES UN ALIAS DE 'uint8'
	var ageBro byte = 25
	fmt.Printf("TIPO: %T, VALOR: %v \n", ageBro, ageBro)

	// ALIAS DE int32 'rune' ES ESCENIAL PARA MANEJAR CORRECTAMENTE LOS CARACTERES UNICODE.
	// CADENAS QUE CONTIENEN CARACTERES FUERA DEL RANGO ASCII ESTANDAR

	var s rune = 'a'

	fmt.Printf("TIPO: %T, VALOR: %v \n", s, s)

	// VALORES FLOTANTES

	var pi float32 = 3.1416
	fmt.Printf("TIPO: %T, VALOR: %v \n", pi, pi)

	//! DADO QUE GO ES ESTATICAMENTE TIPADO... NO DEJA HACER OPERACIONES DIRECTAS ENTRE LOS DIFERNETES TIPOS DE DATOS
	// POR LO TANNTO EL TIPO DE DATO NO PUEDE CAMBIAR

	/*
		var num1 int8 = 127
		var num2 int32 = 10000

		c := num1 + num2

		fmt.Printf("TIPO: %T, VALOR: %v", c, c)
		! ESTO NO SE PUEDE :(
	*/

	// ? PERO PARA HACER OPERACIONES ASI TENEMOS QUE HACER UN CASTING

	var num1 int8 = 127
	var num2 int32 = 10000

	// * PARA HACERLO DEBEMOS OCUPAR LA FUNCION DEL TIPO DE DATO QUE QUEREMOS QUE SER CONVIERTA
	c := int32(num1) + num2

	fmt.Printf("TIPO: %T, VALOR: %v \n", c, c)

	// ? SIN EMBARGO A PESAR DEL CASTING EL TIPO DE DATO ORIGINAL NUNCA VA A CAMBIAR
	fmt.Printf("TIPO: %T, VALOR: %v \n", num1, num1)

	// EN GO EXISTE ALGO LLAMADO VALOR CERO... CUANDO DECALRAS UNA VARIABLE O CONSTANTE GO LE ASIGNA UN VALOR POR DEFECTO QUE ES CERO:

	var zero string
	var zero1 int
	var isbol bool

	fmt.Printf("TIPO: %T, VALOR: %q \n", zero, zero)
	fmt.Printf("TIPO: %T, VALOR: %v \n", zero1, zero1)
	fmt.Printf("TIPO: %T, VALOR: %v \n", isbol, isbol) // SI VALOR CERO SIEMPRE ES 'false'

}
