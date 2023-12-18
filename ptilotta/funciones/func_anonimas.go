package funciones

import "fmt"

func Calculos() {
	// Declaración de una función anónima

	suma := func(a int, b int) int {
		return a + b
	}
	fmt.Println(suma(10, 20))
}
