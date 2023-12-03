package main

import (
	"fmt"

	"github.com/devgarrido/godesde0/ptilotta/variables"
)

func main() {
	variables.ShowIntegers() // Call to ShowIntegers function from variables package.
	variables.OtherVariables()
	fmt.Println(variables.ConvertToText(10))
}
