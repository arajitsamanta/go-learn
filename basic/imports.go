package basic

/*
This code groups the imports into a parenthesized, "factored" import statement.
You can also write multiple import statements, like:

import "fmt"
import "math"

But it is good style to use the factored import statement.
*/
import (
	"fmt"
	"math"
)

//TestImport function
func TestImport() {
	fmt.Printf("****Running gotour1.TestImport(), Now you have %g problems.", math.Sqrt(7))
}
