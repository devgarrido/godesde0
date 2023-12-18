package concept

import "fmt"

func ShowMaps() {
	// Maps is a collection of unordered key-value pairs.
	// Example:
	// music := make(map[string]string) --> music is a map with string keys and string values.
	music := make(map[string]string)
	music["guitar"] = "acoustic"
	music["drums"] = "electronic"
	music["violins"] = "stradivarius"
	fmt.Printf("%+v\n", music)

	// Another way to declare a map:
	// music2 := map[string]string{}
	pets := map[string]string{"dog": "puppy", "cat": "kitten", "fish": "fishling"}
	fmt.Printf("%+v\n", pets)
	fmt.Printf("%+v\n", pets["cat"])

	// Delete a key-value pair from a map:
	delete(pets, "fish")
	fmt.Printf("%+v\n", pets)

	// Check if a key exists in a map:
	_, ok := pets["fish"]
	if ok {
		fmt.Println("The key exists")
	} else {
		fmt.Println("The key doesn't exist")
	}

	// Add a key-value pair to a map:
	pets["fish"] = "fishling"
	fmt.Printf("%+v\n", pets)

	// Iterate over a map:
	for key, value := range pets {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// Iterate over a map without using the key:
	for _, value := range pets {
		fmt.Printf("Value: %s\n", value)
	}

}
