// Un slice en GO es un array dinámico, es decir, que puede cambiar de tamaño.
// La definición de un slice es similar a la de un array, pero sin indicar el tamaño.
// Los slices son apuntadores a un array, por lo que si modificamos un slice, modificamos el array original.
// los slices son una coleccion de elementos del mismo tipo y con un tamaño dinamico.
package slices

import (
	"fmt"
	"reflect"
)

var Pf = fmt.Printf

func ShowSlices() {
	// Creamos un array de 7 elementos de tipo string
	arrayThings := [7]string{"VW Polo", "VW Golf", "VW Passat", "Audi A3", "Audi Q5", "Globo", "Avion"}
	car_VW := arrayThings[0:4] // El slice car apunta a los elementos del array desde el indice 0 hasta el 3 (el 4 no se incluye).
	Pf("array of things: %v\n", arrayThings)
	Pf("array of cars: %v\n", car_VW)
	println("***********************************************************************")
	// Si modificamos el slice car, tambien modificamos el array original
	car_VW[0] = "VW Tiguan"
	Pf("array of things: %v\n", arrayThings)
	car_Audi := arrayThings[3:5]
	Pf("array of cars: %v\n", car_Audi)

	// Los elementos que van [] se llaman indices, y los elementos que van dentro de [] se llaman elementos.
	// [:] -> todos los elementos
	// [2:] -> desde el elemento 2 hasta el final
	// [:5] -> desde el principio hasta el elemento 5 (sin incluirlo)
	// [2:5] -> desde el elemento 2 hasta el elemento 5 (sin incluirlo)

}
func ShowSlices2() {

	// len() -> longitud del slice
	// cap() -> capacidad del slice. Es el numero de elementos que puede contener el slice a partir del indice n
	// append() -> agrega un elemento al slice

	animals := []string{"🦍", "🐕", "🐶", "🐦", "🐘"}
	pets := animals[1:3]
	fmt.Println("animals:", animals)
	fmt.Println("pets:", pets)
	fmt.Println("len(pets):", len(pets))
	fmt.Println("cap(pets):", cap(pets))
	fmt.Println("Tipo de dato de pets:", reflect.TypeOf(pets))

	// En esta caso no existe un array original, el slice es el array.
	var sliceFlags []string
	sliceFlags = append(sliceFlags, "SP") // append agrega un elemento al slice
	sliceFlags = append(sliceFlags, "FR")
	sliceFlags = append(sliceFlags, "IT")

	Pf("Banderas: %v\n", sliceFlags)

	// Tambien podemos crear un slice con un slice literal
	sliceFlags2 := []string{"SP", "FR", "IT"}
	Pf("Banderas: %v\n", sliceFlags2)

	// Otra forma de declarar un slice es con la funcion make
	sliceFlags3 := make([]string, 3)
	sliceFlags3[0] = "SP"
	sliceFlags3[0] = "SP"
	sliceFlags3[1] = "FR"
	sliceFlags3[2] = "IT"
	sliceFlags3 = append(sliceFlags3, "DE") // append agrega un elemento al slice
	Pf("Banderas: %v\n", sliceFlags3)

	// Tambien podemos crear un slice con un slice literal

}
