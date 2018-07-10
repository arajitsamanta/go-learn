package flowcontrol

import (
	"fmt"
	"math"
)

//Sqrt - Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.
func Sqrt(x float64) string {
	fmt.Printf("\n****Running flowcontrol.Sqrt(), go if statement example ")
	if x < 0 {
		return Sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

/*
IfWithShortStatement - Like for, the if statement can start with a short statement to execute before the condition.
Variables declared by the statement are only in scope until the end of the if.

(Try using v in the last return statement.)
*/
func IfWithShortStatement(x, n, lim float64) float64 {
	fmt.Printf("\n****Running flowcontrol.IfWithShortStatement(), if with short statement ")
	if v := math.Pow(x, n); v < lim {
		return v
	}
	// can't use v here
	return lim
}

/*
IfWithElse - Variables declared inside an if short statement are also available inside any of the else blocks.
(Both calls to pow are executed and return before the call to fmt.Println in main begins.)
*/
func IfWithElse(x, n, lim float64) float64 {
	fmt.Printf("\n****Running flowcontrol.IfWithElse(), go if else statement ")
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
