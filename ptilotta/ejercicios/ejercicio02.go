package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func TablaMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)
	var numero int
	var err error
	var resultado string
	for {
		fmt.Println("Ingrese un número: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}
	for i := 1; i <= 10; i++ {
		resultado += fmt.Sprintf("%d x %d = % d\n", numero, i, numero*i)
	}

	return resultado
}
