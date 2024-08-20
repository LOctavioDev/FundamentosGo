package main

import "fmt"

func main() {
	// ? SLICE: SON APUNTADORES A UN ARRAY, COMO TAL NO POSEEN DATOS
	things := [7]string{"🦴", "🦿", "👀", "🦷", "🍫", "🍬", "🍭"}

	// ? LO QUE HACE ES SEPARAR LOS ELEMENTOS DE UN ARREGLO, DESDE UNA POSICION INICIAL A UNA FINAL ALGO COMO ESTO: things[P_INICIAL:P_FINAL+1] LA PRIMERA POSICION ES INCLUYENTE, PERO LA ULTIMA ES DECLUYENTE OSEA NO LA TOMA EN CUENTA, TENDRIAS QUE AUMENTAR UNA POSICION

	whiteThigns := things[0:4] // * 🦴 🦿 👀 🦷
	candys := things[4:7]      // * 🍫 🍬 🍭

	// ! COMO NO ESTAMOS CREANDO UN NUEVO ARRAY, SI MODIFICAMOS UNA POSICION DE UN PUNTERO A ESE ARRAY, SE MODIFICA EL ORIGINAL TAMBIEN
	candys[0] = "🍮" // LO REEMPLAZA, DESDE EL PUNTERO PUEDE MOFICIAR EL ARRAY ORIGINAL

	fmt.Println("Things:", things)
	fmt.Println("White things:", whiteThigns)
	fmt.Println("Candys:", candys)

	// TAMPOCO ES NECESARIO INDICAR SIEMPRE EL INDICE, GO LO INFIERE
	fmt.Println(things[3:])
	fmt.Println("==============================================================================")

	animals := [7]string{"🐒", "🐎", "🐕", "🐈", "🐇", "🦌", "🦬"}
	pets := animals[2:5]

	// ! RECORDANDO QUE EN GO LOS ARRAEGLOS TIENE UN TAMANIO DEFINIDO, SI AGRGAMOS MAS ELEMENTOS Y YA NO SOPORTA LA CANTIDAD, LO QUE HACE GO ES CREAR UN NUEVO ARRAY Y DUPLICA LA CAPACIDAD DEL SLICE ORIGINAL
	pets = append(pets, "🐀", "🐩", "🐢")

	fmt.Println("Animals:", animals)
	fmt.Println("Pets:", pets)

	// HAY UNA FUNCION QUE YA ESTA CREADA EN GOLANG, QUE ES PARA PODER SABER EL TAMANIO DE UN ARRAY "len()"
	// HAY UNA FUNCION PARA SABER EL NUMERO DE ELEMENTOS DEL ARRAY DE ORIGEN, APARTIR DEL INDICE DESD SE CREO EL SLICE

	fmt.Println("Tamanio de pets:", len(pets))
	fmt.Println("Capacidad de pets:", cap(pets))

	fmt.Println("==============================================================================")

	// ? EN GO SE PUEDE CREAR UN SLICE SIN LA NECECIDAD DE QUE EXISTA UN ARRAY, ACA TENEMOS UN EJEMPLO CON EL METODO PRECONSTRUDIO: "make()" DONDE PRIMERO LE DEBEMOS PASAR DE QUE TIPO VA A SER NUESTO SLICE Y DESPUES EL TAMANIO DEL SLICE

	emojis := make([]string, 0, 3)
	emojis = append(emojis, "😀", "😁", "😂", "😊")

	fmt.Println("Emojis:", emojis)
	fmt.Println("Tamanio de emojis:", len(emojis))
	fmt.Println("Capacidad de emojis:", cap(emojis))
}
