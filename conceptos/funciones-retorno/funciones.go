package main

import (
	"fmt"
	"strings"
)

// ? EN GO TAMBIEN LAS FUNCIONES PUEDNE RETORNAR VALORES, SOLO HAY QUE ESPECIFICAR QUE TIPO DE DATO SE QUIERE RETORNAR

func main() {
	result := addNumbers(5, 2)
	fmt.Println(result)

	language := "PythonS"
	v1, v2 := convertString(language)
	fmt.Println(v1, v2)
}

func addNumbers(a, b int) int {
	return a + b
}

// ? PARA QUE LAS FUNCIONES PUEDAN RETORNAR VARIOS VALORES, ENTRE PARENTESIS ESPECIFICAMOS DE QUE TIPO DE DATOS SON
func convertString(a string) (string, string) {
	lower := strings.ToLower(a)
	upper := strings.ToUpper(a)
	return lower, upper
}
