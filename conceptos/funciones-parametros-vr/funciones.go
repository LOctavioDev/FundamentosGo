package main

import (
	"fmt"
	"strings"
)

// ? LAS FUNCIONES EN GO, SON MUY SIMILARES AL ESTANDAR EN LOS DEMAS LENGUAJES. PRIMERO USAMOS LA PALABRA RESEVADA " func " SEGUIDO DEL NOMBRE DE LA FUNCION, POSTERIORMENTE ENTRE PARENTESIS PUEDEN IR O NO LOS PARAMTROS QUE RECIBIRA LA FUNCION. Y ENTRE LLAVES " {} " YA LA LOGICA DE LA FUNCION EN SI.

func main() {
	// PARA INVOCAR UNA FUNCION SOLO TENEMOS QUE USAR EL NOMBRE DE LA FUNCION SEGUIDO DE LOS PARENTESIS

	greet()
	greetWithParams("Jesus", "Entradas")

	name := "haziel"
	toUpperCase(&name)
	fmt.Println(name)
}

// * FUNCION DE SALUDAR
func greet() {
	fmt.Println("HELLO MOTHER FUCKERS")
}

// * FUNCION SALUDAR CON PARAMETROS
func greetWithParams(firtsName, lastName string) {
	fmt.Println("Hello", firtsName, lastName, "fuck your mother")
}

// * FUNCION QUE CONVIERTA UN TEXTO EN MAYUSCULAS
// ! DADO QUE EN GO CUANDO CUANDO PASAMOS UN VALOR A UNA FUNCION NO SE MODIFICA DIRECTAMENTE ESE VALOR, YA QUE POR DEFECTO SE PAS AUNA COPIA DE ESE VALOR. O SEA QUE SI QUEREMOS HACER CAMBIOS DE UN VALOR DENTRO DE UNA FUNCION, ESOS CAMIOS NO AFECTARAN EL VALOR ORIGIAL FUERA DE LA FUNCION
func toUpperCase(text *string)  {
	*text = strings.ToUpper(*text)
}