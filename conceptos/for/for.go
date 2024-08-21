package main

import "fmt"

func main() {
	// ? EN LOS DEMAS LENGUAJES DE PROGRAMACION TENEMOS LAS ESTRUCTURAS DE CONTROL QUE TODOS CONOCEMOS PARA BUCLES, EL WHILE, DO WHILE, FOR. EN ESTE CASO, GO SOLO TIENE EL " for " SIN EMBARGO ESTE TIENE VARIAS FORMAS DE IMPMENETARSE PARA ADAPTARSE A DIREFENTES SITUACIONES.

	// * UNA FORMA ES LA MAS CLASICA, LA QUE TODOS CONOCEMOS PERO COMO HAS NOTADO, NO SE TIENEN QUE PONER ENTRE PARENTESIS LA CONDICION

	// MOSTRAR LA TABLA DE MULTIPLICAR DEL 10
	for i := 1; i <= 10; i++ {
		fmt.Printf("10 X %d = %d\n", i, i*10)
	}

	fmt.Println("==================================================")

	// * PARA TENER UN BUCLE SIMILAR A WHILE SIMPLEMENTE SOLO DEJAS LA CONDICION Y YA :).

	// MOSTRAR LA TABLA DE MULTIPLICAR DEL 2

	i := 1

	for i <= 10 {
		fmt.Printf("2 X %d = %d\n", i, i*2)
		i++
	}

	fmt.Println("==================================================")

	// * PARA TENER UN BUCLE INDEFINIDAMENTE Y QUE SE EJECUTE SIEMPRE SOLO PONEMOS " for{} " ASI CONDICION :).

	// ! ESTO ES UTIL EN CIERTOS CASOS CUANDO TIENE QUE ESTAR SIEMPRE EN LA ESCUCHA
	// ! SI QUEREMOS ROMPER EL CICLO USAMOS LA PALABRA RESERVADA " breake "

	j := 0

	for {
		if j == 10 {
			break // SI SE CUMPLE PUES QUE TRUENE :3
		}

		fmt.Println("SIEMPRE, PENDENTES ;) QUE EL GOBIERNO ES MUY INTELIGENTE")

		j++
	}

	fmt.Println("==================================================")

	numbers := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// ? PARA RECORRER UN SLICE EN GO PODEMOS USAR LA PALABRA RESERVADA " range " SEGUIDO DEL ITERABLE
	// RANGE NOS DEVUELVE DOS VALORES, EL PRIMERO ES EL INDICE, Y EL SEGUNDO ES EL VALOR
	for k, v := range numbers {
		fmt.Println("INDICE: ", k, "VALOR:", v)
	}

	fmt.Println("==================================================")

	numbers2 := []uint8{1, 2, 3, 4, 5}

	for i := range numbers2 {
		numbers2[i] *= 2
	}

	fmt.Println(numbers2)

	fmt.Println("==================================================")

	// ? ES LA MISMA LOGICA PARA RECORRER UN MAPA

	food := map[string]string{
		"pizza":  "🍕",
		"burger": "🍔",
		"hocho":  "🌭",
		"soda":   "🥤",
	}

	for key, value := range food {
		fmt.Println("KEY:", key, "VALUE:", value)
	}

	fmt.Println("==================================================")

	// ? TAMBIEN SE PUEDE ITERAR UN STRING

	for _, v := range "aeiou"{
		fmt.Println(string(v))
	}


}
