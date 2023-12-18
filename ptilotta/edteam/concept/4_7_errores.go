// Los errores en Go son un tipo de dato, por lo que se pueden asignar a una variable
// y se pueden retornar como valor de una funcion.
// Los errores se pueden crear con la funcion errors.New("mensaje de error")

// Los errores se pueden crear con la funcion fmt.Errorf("mensaje de error")
// Los errores se pueden crear con la funcion fmt.Sprintf("mensaje de error")
// Los errores se pueden crear con la funcion fmt.Errorf("mensaje de error").

package concept

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

var errNotFound = errors.New("not found error")

// Foods map
var foods = map[int]string{
	1: "🍔",
	2: "🍕",
	3: "🌭",
	4: "🍿",
	5: "🍩",
}

func ShowErrors() {
	value, err := strconv.Atoi("33")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(value)

	found, err := search("56")
	if errors.Is(err, errNotFound) {
		fmt.Println("Error controlado por nosotros")
		return
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(found)
}

func search(key string) (string, error) {
	num, err := strconv.Atoi(key)
	if err != nil {
		return "", fmt.Errorf("strconv.Atoi(key) %w", err)
	}
	emoji, err := findFood(num)
	if err != nil {
		return "", fmt.Errorf("findFood(num) %w", err)
	}
	return emoji, nil
}

func findFood(id int) (string, error) {
	food, ok := foods[id]
	if !ok {
		return "", errNotFound
	}
	return food, nil
}
