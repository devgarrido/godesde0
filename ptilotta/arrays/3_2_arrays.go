package arrays

import "fmt"

var Pf = fmt.Printf

func ShowArrays() {
	// los arrays son una coleccion de elementos del mismo tipo y con un tamaño fijo.
	var arrayFlags [3]string
	arrayFlags[0] = "SP"
	arrayFlags[1] = "FR"
	arrayFlags[2] = "IT"
	Pf("Banderas: %v\n", arrayFlags)

	// Mostramos los arrays con ciclo for
	for i := 0; i < len(arrayFlags); i++ {
		Pf("Banderas: %v\n", arrayFlags[i])
	}

	// Mostramos los arrays con ciclo for range (mas eficiente), Muestra el indice y el valor.
	for i, v := range arrayFlags {
		Pf("Banderas: %v - %v\n", i, v)
	}

	// Mostramos los arrays con ciclo for range (mas eficiente), Muestra solo el valor.
	for _, v := range arrayFlags {
		Pf("Banderas: %v\n", v)
	}

	// Array literal
	arrayFruits := [3]string{"Apple", "Orange", "Banana"}
	for _, v := range arrayFruits {
		Pf("Fruits: %v\n", v)
	}

	// Array con inferencia de tamaño
	arrayFruits2 := [...]string{"Apple", "Orange", "Banana"}
	for i := 0; i < len(arrayFruits2); i++ {
		Pf("Fruits: %v\n", arrayFruits2[i])
	}

}
