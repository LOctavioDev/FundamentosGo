package main

import "fmt"

func main() {
	// ? SLICE: SON APUNTADORES A UN ARRAY, COMO TAL NO POSEEN DATOS
	things := [7]string{"🦴", "🦿", "👀", "🦷", "🍫", "🍬", "🍭"}

	// ? LO QUE HACE ES SEPARAR LOS ELEMENTOS DE UN ARREGLO, DESDE UNA POSICION INICIAL A UNA FINAL ALGO COMO ESTO: things[P_INICIAL:P_FINAL+1] LA PRIMERA POSICION ES INCLUYENTE, PERO LA ULTIMA ES DECLUYENTE OSEA NO LA TOMA EN CUENTA, TENDRIAS QUE AUMENTAR UNA POSICION

	whiteThigns := things[0:4] // 🦴 🦿 👀 🦷
	candys := things[4:7]      // 🍫 🍬 🍭

	// ! COMO NO ESTAMOS CREANDO UN NUEVO ARRAY, SI MODIFICAMOS UNA POSICION DE UN PUNTERO A ESE ARRAY, SE MODIFICA EL ORIGINAL TAMBIEN
	candys[0] = "🍮" // LO REEMPLAZA, DESDE EL PUNTERO PUEDE MOFICIAR EL ARRAY ORIGINAL

	fmt.Println("Things: ", things)
	fmt.Println("White things: ", whiteThigns)
	fmt.Println("Candys: ", candys)

	// TAMPOCO ES NECESARIO INDICAR SIEMPRE EL INDICE, GO LO INFIERE
	fmt.Println(things[3:])

}
