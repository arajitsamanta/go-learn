package gotour1

import "fmt"

var c, python, java bool

/*
VaribalesDeclaration - example

The var statement declares a list of variables; as in function argument lists, the type is last.
A var statement can be at package or function level. We see both in this example.
*/
func VaribalesDeclaration() {
	fmt.Printf("\n****Running gotour1.VaribalesDeclaration(), variable declaration example  ")
	var i int
	fmt.Println(i, c, python, java)
}

var i, j int = 1, 2

/*
VaribalesWithInitializers  - example

A var declaration can include initializers, one per variable.
If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
*/
func VaribalesWithInitializers() {
	fmt.Printf("****Running gotour1.VaribalesWithInitializers(), varbiable initialization values --> ")
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

/*
ShortVaribaleDeclaration  - example

Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
*/
func ShortVaribaleDeclaration() {
	fmt.Printf("****Running gotour1.ShortVaribaleDeclaration(), short variable assignment initialization values --> ")
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}

/*
ZeroValueVariable - example

Variables declared without an explicit initial value are given their zero value.
The zero value is:
0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.
*/
func ZeroValueVariable() {
	fmt.Print("****Running gotour1.ZeroValueVariable(), zero value variables values --> ")
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
