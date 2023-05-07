package main

import (
	"fmt"

	"github.com/devgarrido/godesde0/variables"
)

func main() {
	var1, var2 := variables.ConviertoaTexto(1209)
	if var1 {
		fmt.Println("El valor es\n" + var2)
	}

}
