package concept

import "fmt"

func ShowFuncFunc() {
	nums := []int{43, 20, 12, 9, 20, 78, 11, 29, 52}

	// filtrar numeros pares.
	resultPares := filter(nums, func(n int) bool {
		return n%2 == 0
	})

	// filtrar numeros impares.
	resultNones := filter(nums, numbersOdd)

	// Ver el resultado de la funcion filter.
	for _, n := range resultPares {
		println(n)
	}

	for _, n := range resultNones {
		println(n)
	}

	fmt.Println(sum(2)(3))  // 5
	fmt.Println(sum(5)(5))  // 10
	fmt.Println(sum(10)(2)) // 12

	plus10 := sum(39)
	fmt.Println(plus10(5)) // 44
	fmt.Println(plus10(2)) // 41
	fmt.Println(plus10(1)) // 40

}
func numbersOdd(n int) bool {
	return n%2 != 0
}

// Funcion que recibe otra funcion como parametro.
func filter(numbers []int, callback func(int) bool) []int {
	result := make([]int, 0, len(numbers))
	for _, n := range numbers {
		if callback(n) {
			result = append(result, n)
		}
	}
	return result
}

// Funcion que retorna otra funcion.
// Vamos a realizar una implementación de sumar 2 numeros.

// 1. Definir la funcion que retorna otra funcion.
func sum(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
