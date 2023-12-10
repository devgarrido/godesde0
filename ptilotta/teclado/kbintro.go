package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1, num2 int
var message string
var err error

func IngresoNumeros() {
	scanner := bufio.NewScanner(os.Stdin) // To read from keyboard

	fmt.Println("Ingrese numero-1:")
	if scanner.Scan() {
		num1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Ha ocurrido un error" + err.Error())
		}
	}
	fmt.Println("Ingrese numero-2:")
	if scanner.Scan() {
		num2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Ha ocurrido un error" + err.Error())
		}
	}
	fmt.Println("Ingrese Mensaje:")
	if scanner.Scan() {
		message = scanner.Text()
	}

	fmt.Println(message, num1*num2)
}
