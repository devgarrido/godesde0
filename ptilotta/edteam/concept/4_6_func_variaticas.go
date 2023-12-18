package concept

func ShowFuncVariaticas() {
	valor := sumar(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	println(valor)
	valor = sumar(2, 4, 6, 8, 10)
	println(valor)

	// FUNCIONES ANONIMAS
	// Son funciones que no tienen nombre.
	// Se utilizan para realizar tareas rapidas.

	func() {
		println("Hola Antonio Garrido, eres un máquina")
	}() // () para ejecutar la funcion.

	func(a, b int) {
		println(a + b)
	}(4, 2)

	// Other example.
	greet := func() {
		println("Hola mundo")
	}
	greet()
}

// Funciones variaticas.
// Son funciones que reciben un numero indeterminado de parametros.
func sumar(nums ...int) (result int) {
	for _, n := range nums {
		result += n
	}
	return
}
