// Panic es una palabra reservada que se utiliza para detener la ejecucion de una funcion.
// Se ejecuta en el punto donde se produce el error.
package concept

import (
	"fmt"
)

var pf = fmt.Printf

func ShowPanicRecover() {
	pf("El resultado de la division es: %d\n", division(100, 10)) // 10
	pf("El resultado de la division es: %d\n", division(200, 25)) // 8
	pf("El resultado de la division es: %d\n", division(34, 0))   // 0 // panic: runtime error: integer divide by zero
	pf("El resultado de la division es: %d\n", division(124, 8))  // 15
	pf("El resultado de la division es: %d\n", division(380, 20)) // 19
}

func division(dividendo, divisor int) int {
	// utilizamos defer en una funcion anonima para recuperar el control de la funcion.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperando el control de la funcion")
		}
	}()
	validateZero(divisor)
	return dividendo / divisor
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("No se puede dividir por cero")
	}

}
