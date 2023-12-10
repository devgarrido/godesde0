package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var pf = fmt.Printf

func IntroNum() {
	scanner := bufio.NewScanner(os.Stdin)
	println("Ingrese un número: ")
	scanner.Scan()
	if scanner.Err() != nil {
		// handle error.
		panic("Ha ocurrido un error" + scanner.Err().Error())
	}
	num, _ := strconv.Atoi(scanner.Text())
	// Tabla de multicplicación.
	for i := 1; i <= 10; i++ {
		println(num, "x", i, "=", num*i)
	}
}

func TablaMultiplicar() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un número: ")
		if scanner.Scan() {
			numero, err := strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				for i := 1; i <= 10; i++ {
					pf("numero %d * %d, es = % d\n", numero, i, numero*i)
				}
				break
			}
		}

	}
}
