package concept

import "fmt"

func ShowIfSwitch() {
	/*
		if condition {
			code
		} else if condition {
			code
		} else {
			code
		}
	*/

	if superheroe := "batman"; superheroe == "superman" {
		fmt.Println("Superman")
	} else if superheroe == "batman" {
		fmt.Println("Batman")
	} else {
		fmt.Println("No es un superheroe")
	}

	// whit switch
	switch superheroe := "batman"; superheroe {
	case "superman":
		fmt.Println("Superman")
	case "batman":
		fmt.Println("Batman")
	default:
		fmt.Println("No es un superheroe")
	}

	// whit switch grouping cases
	switch superheroe := "batman"; superheroe {
	case "superman", "batman":
		fmt.Println("Es un superheroe")
	default:
		fmt.Println("No es un superheroe")
	}

	// whit switch without condition
	superheroe := "batman"
	canSearch := false
	switch {
	case !canSearch:
		fmt.Println("No se puede buscar")
		// fallthrough
	case superheroe == "superman", superheroe == "batman":
		fmt.Println("Es un superheroe!!!")
	default:
		fmt.Println("No es un superheroe")
	}
}
