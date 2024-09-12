package main

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func main() {

	// ? INSTANCIA CON LOS NOMBRES DE LOS CAMPOS
	Go := Course{
		Name:    "Go desde Cero",
		Price:   12.34,
		IsFree:  false,
		UserIDs: []uint{12, 56, 89},
		Classes: map[uint]string{
			1: "Introduccion",
			2: "Estructuras",
			3: "Mapas",
		},
	}

	// ? INSTANCIA SIN LOS NOMBRES DE LOS CAMPOS, PERO PARA QUE PUEDA FUNCIONAR TIENEN QUE ESTAR EN EL MISMO ORDEN Y TIENEN QUE ESTAR TODOS
	JavaScript := Course{
		"JavaScript desde Cero",
		15.90,
		false,
		[]uint{13, 56, 78},
		map[uint]string{
			1: "Introduccion",
			2: "DOM",
			3: "Funciones",
		},
	}

	// ? IGUAL PODEMOS ASIGNAR SOLO UNOS CAMPOS
	CSS := Course{Name: "CSS desde Cero", IsFree: true}

	// ? PODEMOS CREAR LA INSTANCIA SON ASIGNAR NADA Y HACERLO DESPUES CON EL "."
	Python := Course{}
	Python.Name = "Python desde Cero"
	Python.UserIDs = []uint{44, 66, 77}

	fmt.Println(Go.Name)
	fmt.Println(JavaScript.Name)
	fmt.Printf("%+v\n", CSS)
	fmt.Printf("%+v", Python)

}
