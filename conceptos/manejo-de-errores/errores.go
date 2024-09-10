package main

import (
	"errors"
	"fmt"
	"strconv"
)

// ACA CREAMOS UN ERROR CON EL PAQUETE DE ERRORS DE GOLANG
var errNotFound = errors.New("not found")

// ? EN GO LOS ERROES SE MANEJAN DE UNA MANERA MUY DIFERENTE A LOS LENGUAJES COMO JAVASCRIPT O PYHTON, DEBIDO A QUE AHI MANEJAS ALGO LLAMADO "try-catch". PERO EN GO SE LLAMA ERRORES COMO TAL, Y SE MANEJAN DE FORMA EXPLICITA.

var food = map[int]string{
	1: "🍕",
	2: "🍔",
}

func main() {
	num, err := strconv.Atoi("1") // * ESA FUNCION HACE QUE CONVIERA UN NUMERO EN CADENA ENTERO

	// PARA VERIFICAR SI HUBO UN ERROR GENERALMENTE SE GUARDA COMO TAL EN UNA VARIABLE Y SI EL VALOR ES DISTINTO DE "nil" QUIERE DECIR QUE SI HAY ERROR. Y DAMOS RETURN PARA QUE NO HAYA PROBLEMAS.
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(num)

	found, errorsito := search("12")

	if errors.Is(errorsito, errNotFound) {
		fmt.Println("Pudimos controlar el error correctamente")
	}

	if errorsito != nil {
		fmt.Println(errorsito)
		return
	}

	fmt.Println(found)

}

func search(key string) (string, error) {
	num, err := strconv.Atoi(key)
	if err != nil {
		return "", fmt.Errorf("strconv.Atoi(): %w", err)
	}

	emoji, err := findFood(num)
	if err != nil {
		return "", fmt.Errorf("findFood(): %w", err)
	}

	return emoji, nil

}

func findFood(id int) (string, error) {
	value, ok := food[id]

	if !ok {
		return "", errNotFound
	}

	return value, nil // ! TAMBIEN RETORNAMOS NIL PORQUE NO HAY NINGUN ERROR.
}
