package basic

import (
	"fmt"
	"math/cmplx"
)

/*
Go's basic types are

bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32, represents a Unicode code point
float32 float64
complex64 complex128

The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.
The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.
*/
var (
	ToBe          = false
	MaxInt uint64 = 1<<64 - 1
	z             = cmplx.Sqrt(-5 + 12i)
)

//DataTypes - Example of Go's basi data types.
func DataTypes() {
	fmt.Print("****Running basic.BasicTypes(), Go's basic data types example")
	fmt.Printf("\nType: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
