package basic

/*
A function can take zero or more arguments.

In this example, add takes two parameters of type int.

Notice that the type comes after the variable name.

(For more about why types look the way they do, see the article on https://blog.golang.org/gos-declaration-syntax)
*/
import "fmt"

//Add - custom function to add two integers
func Add(x int, y int) int {
	sum := x + y
	fmt.Printf("\n****Running gotour1.Add(), sum = %d", sum)
	return sum
}

/*
AddOmitTypeDeclaration - When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

In this example, we shortened
	x int, y int
to
	x, y int
*/
func AddOmitTypeDeclaration(x, y int) int {
	var sum = x + y
	fmt.Printf("\n****Running gotour1.AddOmitTypeDeclaration(), sum = %d", sum)
	return sum
}

//Swap - multiple returns
func Swap(x, y string) (string, string) {
	fmt.Printf("\n****Running gotour1.Swap(), multiple returns")
	return y, x
}
