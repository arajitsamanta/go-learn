package gotour1

import (
	"fmt"
	"math"
)

/*
ConvertTypes - The expression T(v) converts the value v to the type T.

Some numeric conversions:
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
Or, put more simply:
i := 42
f := float64(i)
u := uint(f)

Unlike in C, in Go assignment between items of different type requires an explicit conversion. Try removing the float64 or uint conversions
in the example and see what happens.
*/
func ConvertTypes() {
	fmt.Print("****Running gotour1.ConvertTypes(), type conversions --> ")
	var x, y int = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z = uint(f)
	fmt.Println(x, y, z)
}
