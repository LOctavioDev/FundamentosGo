package main

import "fmt"

type Person struct {
	Name        string
	Age         uint8
	Nacionality string
}

func main() {
	// ? EN GO LA SINTAXIS DEL "if" ES MUY SIMILAR A OTROS LENGUAJES DE PROGRAMACION. SE HACE HACIENDO USO DE LA PALABRA RESERVADA "if" SEGUIDO DE LA CONDICIONAL, SIN EMBARGO, NO SE TIENE QUE PONER ENTRE PARENTESIS LA CONDICION.

	octavio := Person{
		Name:        "Octavio",
		Age:         22,
		Nacionality: "Mexican",
	}

	if octavio.Age <= 22 && octavio.Nacionality != "American" {
		fmt.Println("Draft NBA:", octavio.Name)
	} else if octavio.Age <= 22 && octavio.Nacionality == "American" {
		fmt.Println("Draft NBA, but you are American")
	} else {
		fmt.Println("No draft NBA:", octavio.Name)
	}

	// ?  EN EL CASO DEL SWITCH LA SINTAXIS ES MUY SIMILAR A CASI TODOS LOS LENGAUJES ESTANDAR. CON UNA EXPRESION, UNCA CONDICION Y UN CASO DEFAULT. EN GO NO ES NECESARIO USAR EL BREAK

	switch octavio.Age {
	case 19:
		fmt.Println("You are on time to enter the NBA")
	case 20:
		fmt.Println("You're still in time to enter the NBA")
	case 21:
		fmt.Println("You are about to enter the NBA")
	case 22:
		fmt.Println("This is yout last chance to enter the NBA")
	default:
		fmt.Println("Your can't get into the NBA")
	}

}
