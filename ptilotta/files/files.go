package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/devgarrido/godesde0/ptilotta/ejercicios"
)

const filename = "./ptilotta/files/txt/tabla.txt"

func SaveTable() {
	resultado := ejercicios.TablaMultiplicar()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Ha ocurrido un error al crear el archivo: ", err.Error())
		return
	}
	fmt.Fprintln(archivo, resultado)
	archivo.Close()
}

func AddTable() {
	resultado := ejercicios.TablaMultiplicar()
	if !Append(filename, resultado) {
		fmt.Println("Ha ocurrido un error al concatenar el archivo")
	}

}

func Append(archivo string, texto string) bool {
	arch, err := os.Open(archivo)
	if err != nil {
		fmt.Println("Ha ocurrido un error al abrir el archivo: ", err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Ha ocurrido un error al escribir el archivo: ", err.Error())
		return false
	}
	arch.Close()
	return true
}

func ReadFile() {
	arch, err := os.Open(filename) // For read access.
	if err != nil {
		fmt.Println("Ha ocurrido un error al abrir el archivo: ", err.Error())
		return
	}

	scanner := bufio.NewScanner(arch)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println(">" + registro)
	}
	arch.Close()
}
