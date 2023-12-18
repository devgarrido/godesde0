package concept

import "fmt"

func ShowFor() {

	// 1. for condition {
	for i := 0; i <= 5; i++ {
		fmt.Println("value: ", i)
	}

	// 2. for condition {
	i := 0
	for i <= 5 {
		fmt.Println("value: ", i)
		i++
	}

	// 3. for {
	i = 0
	for {
		fmt.Println("value: ", i)
		i++
		if i > 5 {
			break
		}
	}

	// Iteration in a slice.
	var names = []string{"Antonio", "Garrido", "Carranza"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for i, v := range []string{"pizza", "hamburguesa", "tacos", "tortas"} {
		fmt.Println("index: ", i, "value: ", v)
	}

	// Cambiar el valor de un slice con range

	numbers := []uint8{2, 4, 6, 8, 10}
	for i := range numbers {
		numbers[i] *= 2
	}
	fmt.Println(numbers)

	// Iteration in a map. with range
	food := map[string]string{
		"pizza":     "🍕",
		"hamburger": "🍔",
		"apple":     "🍎",
		"hotdog":    "🌭",
	}

	for i, v := range food {
		fmt.Printf("%s %s\n", i, v) // v returns the value of the key (rune) in the map.
	}

	// Iteracion sobre un string.
	for i, v := range "Hola mundo" {
		fmt.Println(i, string(v))
	}

	// for condition {
	// 	code
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
}
