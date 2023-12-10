package iteraciones

import (
	"fmt"
)

var cars = []string{"VW Polo", "VW Golf", "VW Passat", "Audi A3", "Audi Q5", "Globo", "Avion"}

func Iterar() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// utilizamos el Range para iterar sobre un slice
	for i, value := range cars {
		fmt.Println(i, value)
	}
}
