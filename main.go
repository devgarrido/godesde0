package main

import (
	"fmt"

	"github.com/devgarrido/godesde0/ptilotta/exercises"
	"github.com/devgarrido/godesde0/ptilotta/variables"
)

func main() {

	fmt.Println(exercises.ConverIntegerToString("100"))
	fmt.Println("**************************************************")

	variables.ShowIntegers() // Call to ShowIntegers function from variables package.
	fmt.Println("**************************************************")

	variables.OtherVariables()
	fmt.Println("**************************************************")

	variables.EdVariables()
	fmt.Println("**************************************************")

	fmt.Println(variables.ConvertToText(10))
	fmt.Println("**************************************************")

	variables.ShowConstants()

}
