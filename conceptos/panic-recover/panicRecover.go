package main

import "fmt"

func main() {
	// ? EN CORTAS PALABRAS "panic" SE USA PARA INDICAR ERRORES GRAVES, DETENIENDO EL PROGRAMA JUSTO AHI DONDE DE DETECTO EL ERROR.
	// ? Y PORT SU PARTE "recover" ES LO QUE PERMITE MANEJAR Y EVITAR QUE EL PROGRAMA SE DETENGA COMPLEAMENTE, SOLO SI SE USA DENTRO DE FUNCIONES DIFERIDAS.

	division(100, 10)
	division(46, 4)
	division(20, 0)
	division(50, 3)
}

func division(dividendo, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ya andamos bien")
		}
	}()
	validateZero(divisor)
	fmt.Println(dividendo / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("🤕 no mames como vas a dividir entre 0")
	}
}
