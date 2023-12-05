package exercises

import "strconv"

func ConverIntegerToString(text string) (int, string) {
	var message string
	value_n, err := strconv.Atoi(text)
	if err != nil {
		return 0, "La has cagado\n" + err.Error()
	}
	if value_n > 100 {
		message = "The value is > 100"
	} else {
		message = "The value is < 100, gilipollas"
	}
	return value_n, message

}
