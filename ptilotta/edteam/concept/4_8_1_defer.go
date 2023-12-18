// Defer es una palabra reservada que se utiliza para posponer la ejecucion de una funcion o linea de codigo.
// Entra en una pila de ejecucion y se ejecuta al final de la funcion.

package concept

import (
	"fmt"
	"os"
)

func ShowDefer() {

	// Esta función creara un archivo y escribira en el "Hola Antonio Garrido, eres un máquina".
	// Si se produce un error, se mostrara por pantalla.
	// Si no se produce un error, se cerrara el archivo.
	// Para ver el contenido del archivo, ejecutar en la terminal: cat test.txt
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close() // Cuando se termine de ejecutar la funcion, se ejecutara esta linea.(defer = posponer).

	_, err = file.Write([]byte("Hola Antonio Garrido, eres un máquina"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
