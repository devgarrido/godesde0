// Los punteros son variables que almacenan direcciones de memoria de otras variables.
//
// Para ver la direccion de memoria de una variable, usamos el operador &:
// fmt.Println(&variable). Notar que el tipo de dato de la direccion de memoria es *int.
//
// Para declarar un puntero, usamos el operador *: var puntero *int.

package punteros

import "fmt"

var Pf = fmt.Printf

func ShowPunteros() {
	var color string = "🧧"
	// declaración de puntero
	var pointerColor *string = &color // declaracion y asignación de puntero

	Pf("Tipo: %T, Valor: %v, Dirección de memoria: %v\n", color, color, &color)
	Pf("Tipo: %T, Desreferenciación: %s, Dirección de memoria: %v\n", pointerColor, *pointerColor, pointerColor)

	// Cambiamos el valor de la variable color a través del puntero
	*pointerColor = "🎁"
	Pf("Tipo de dato: %T, Valor desreferenciado: %s, Dirección de memoria: %v\n", pointerColor, *pointerColor, pointerColor)
	Pf("Tipo de dato: %T, Valor de la variable : %v, Dirección de memoria: %v\n", color, color, &color)
}
