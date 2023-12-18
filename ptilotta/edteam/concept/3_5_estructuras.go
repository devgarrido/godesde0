package concept

import "fmt"

type Person struct { // the struct is a collection of properties. It's like a class in other languages.
	Name        string
	Age         uint8
	HasChildren bool
}

func ShowStructs() {
	alexys := Person{
		Name:        "Alexys",
		Age:         30,
		HasChildren: false,
	}

	// Another way to declare a struct:
	aGarrido := Person{"Antonio Garrido", 58, false}
	aGarrido.Name = "Antonio Garrido Garrido" // Change the value of a property

	// Another way to declare a struct:
	var mCarranza Person
	mCarranza.Name = "Mercedes Carranza"
	mCarranza.Age = 53
	mCarranza.HasChildren = true

	// %+v muestra el nombre de la propiedad
	fmt.Printf("%+v\n", alexys)
	fmt.Printf("%+v\n", aGarrido)
	fmt.Printf("%+v\n", mCarranza)

	cCristina := &mCarranza // cCristina is a pointer to mCarranza
	cCristina.Name = "Cristina Carranza"

	// Change the value of a property using a pointer to the struct
	(*cCristina).Age = 25 // Esta sisntaxis es equivalente a la siguiente:
	cCristina.Age = 26

	fmt.Printf("%+v\n", mCarranza)
	fmt.Printf("%+v\n", cCristina)

}
