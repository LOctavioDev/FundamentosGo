package main

import "fmt"

func main() {
	// ? LOS MAPAS EN GO SON ESTRUCTURAS CLAVE-VALOR, MISMA FILOSOFIA DE UN JSON POR EJEMPLO
	// SE DECLARA USANDO LA PALABRA RESERVADA "map" Y SE INICIALIZA CON LA FUNCION DE "make()"

	music := make(map[string]string)
	music["guitar"] = "🎸"
	music["violin"] = "🎻"

	fmt.Println(music)

	// ? CLARO TAMBIEN PODEMOS DECLARAR Y ASIGNAR AL MISMO TIEMPO

	user := map[string]string{
		"last_name": "Martinez", 
		"name":      "Octavio", // ! EN GO SI ES OBLIGATORIO QUE TODOS LLEVEN " , " AL FINAL
	}
	fmt.Println(user)

	// ? PARA ELIMINAR UN ELEMENTO DE UN MAPA USAMOS LA FUNCION DE " delete() " PASANDO COMO PARAMETRO EL MAPA Y DESPUES LA CLAVE A ELIMINAR

	delete(user, "name")

	fmt.Println(user)

	// ? PARA PODER LEER PROPIEDADES DE UN MAP, ES COMO EN LOS LENGUAJES DE PROGRAMACION CLASICOS:

	fmt.Println(music["guitar"]) // SI LA CLAVE NO EXISTE VERAS UN VACIO

	// ? EN LAS FUCNIONES DE GO, PUEDES DEVOLVER VARIAS COSAS. POR EJEMPLO CUANDO ACCEDEMOS A UN PROPIEDAD DE MAP, NOS RETORNA EL CONTENIDO EN SI, Y UN VALOR BOOL QUE SI EXISTE ESA CLAVE O NO

	content, isExist := music["violin"]
	// content, isExist := music["MarianoYRomero"]

	fmt.Println(content, isExist)

}
