package ejercicios

import "strconv"

func Grefusa(value string) (int, string) {
	var message string
	text, err := strconv.Atoi(value)
	if err != nil {
		return 0, "La has cagado"

	}
	if text > 100 {
		message = "> 100"
	} else {
		message = "< 100, gilipollas"
	}
	return text, message
}
