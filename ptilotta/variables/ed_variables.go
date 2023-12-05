package variables

import "fmt"

func EdVariables() {
	var apple, banana, orange string = "banana", "apple", "orange" // declare multiple variables of the same type
	var (
		one   int
		two   int
		three int // declare multiple variables of the same type
	)
	one, two, three = 1, 2, 3 // assign values to multiple variables of the same type
	fmt.Printf("apple: %s\n", apple)
	fmt.Printf("banana: %s\n", banana)
	fmt.Printf("orange: %s\n", orange)
	fmt.Printf("one: %d\n", one)
	fmt.Printf("two: %d\n", two)
	fmt.Printf("three: %d\n", three)

	// Declare multiple variables of different types
	var (
		name   string  = "Antonio"
		age    int     = 58
		salary float32 = 1820.89
	)

	// Verbs for formatting
	// %s - string
	// %d - decimal
	// %f - float
	// %t - boolean
	// %v - any value
	// %T - type of the value
	// %p - pointer

	fmt.Printf("name: %s age: %d salary: %f\n", name, age, salary)
}
