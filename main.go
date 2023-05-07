package main

import (
	"fmt"
	"runtime"

	"github.com/devgarrido/godesde0/ejercicios"
	"github.com/devgarrido/godesde0/variables"
)

func main() {
	var1, var2 := variables.ConviertoaTexto(1209)
	if !var1 {
		fmt.Println("El valor es\n" + var2)
	}

	if os := runtime.GOOS; os != "windows" {
		fmt.Println("Este sistema no es windows")
	} else {
		fmt.Println("Este sistema es windows")
	}

	// Vamos a probar el switch

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Esto es windows")
	case "linux":
		fmt.Println("Esto es linux")
	default:
		fmt.Printf("%s\n", os)
	}

	fmt.Println("********************************")
	valor1, valor2 := ejercicios.Grefusa("10asd")
	fmt.Println(valor1, valor2)
}
