package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Name string
var Status bool
var Salary float32
var Date time.Time

func OtherVariables() {
	Name = "Antonio"
	Status = true
	Salary = 1820.89
	Date = time.Now()

	fmt.Println(Name)
	fmt.Println(Status)
	fmt.Println(Salary)
	fmt.Println(Date)
}

func ConvertToText(num int) (bool, string) {
	value := strconv.Itoa(num)
	return true, value
}
