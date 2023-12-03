package variables

import "fmt"

var pf = fmt.Printf

func ShowIntegers() {
	// Declare variables
	var intComun int
	intde32 := int32(10)  // intde32 is an int32 variable
	intde64 := int64(100) // intde64 is an int64 variable

	pf("intComun: %d\n", intComun)
	pf("intde32: %d\n", intde32)
	pf("intde64: %d\n", intde64)
}
