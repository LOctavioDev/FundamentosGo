package main

import "fmt"

// ? RECORDANDO QUE LAS FUNCIONES TAMBIEN SON TIPOS DE DATOS, LAS PODEMOS USAR COMO PARAMETROS DE FUNCION O RETORNOS DE FUNCION

func main() {
	nums := []int{2, 90, 12, 5, 23, 65}
	result := filter(nums, greatherToTen)
	fmt.Println(result)

	result2 := sum(2)(3)
	fmt.Println(result2)

	example := sum(10)

	fmt.Println(example(1))
	fmt.Println(example(2))
	fmt.Println(example(3))
	fmt.Println(example(4))
}

func greatherToTen(num int) bool {
	return num > 10
}

// * FUNCIONES QUE RECIBEN FUNCIONES
func filter(nums []int, callback func(int) bool) []int {
	result := make([]int, 0, len(nums))

	for _, num := range nums {
		if callback(num) {
			result = append(result, num)
		}
	}

	return result
}

// ? FUNCION QUE RETORNA OTRA FUNCION

func sum(a int) func (int) int {
	return func(i int) int {
		return a + i
	}
}