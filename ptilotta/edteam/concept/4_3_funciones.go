package concept

import (
	"fmt"
	"strings"
)

func ShowFunc() {

	name := "antonio garrido garrido"
	toUpperCase(&name)
	fmt.Println(name)

	// suma.
	fmt.Println(suma(10, 20)) // 30.
	upper, lower := convert("Hola Mundo")
	fmt.Println(upper, lower) // hola mundo HOLA MUNDO.

	upper2, lower2 := convert2("Garrido.pro")
	fmt.Println(upper2, lower2) // GARRIDO.PRO garrido.pro.
}

// Esta funcion cambia el valor de la variable name.
// Se han creado 2 variables, una de tipo string y otra de tipo puntero a string.
// Y ambas variables apuntan a la misma direccion de memoria.

func toUpperCase(text *string) {
	*text = strings.ToUpper(*text)
}

// Funciones que retornan valores.
func suma(a, b int) int {
	return a + b
}

// Funciones que retornan varios valores.
func convert(text string) (string, string) {
	return strings.ToUpper(text), strings.ToLower(text)
}

// Funciones que retornan varios valores con nombres.
func convert2(text string) (upper string, lower string) {
	upper = strings.ToUpper(text)
	lower = strings.ToLower(text)
	return
}
