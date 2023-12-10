package ejercicios

import "strconv"

func ConvertText(value string) (int, string) {
	var message string
	text, err := strconv.Atoi(value)
	if err != nil {
		return 0, "La has cagado\n" + err.Error()

	}
	if text > 100 {
		message = "The value is > 100"
	} else {
		message = "The value is < 100, gilipollas"
	}
	return text, message
}
