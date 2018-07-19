package method

import (
	"fmt"
	"math"
)

type vertex3 struct {
	X, Y float64
}

func (v vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
HelloGoMethods - Go does not have classes. However, you can define methods on types.

Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.The receiver appears in its own argument list between the func keyword and
the method name. In this example, the Abs method has a receiver of type Vertex named v.
*/
func HelloGoMethods() {
	fmt.Printf("****Running methods.HelloGoMethods(), go methods \n")
	v := vertex3{3, 4}
	fmt.Println(v.Abs())
}

func abs(v vertex3) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
GoMethodsAreFunctions - a method is just a function with a receiver argument
Remember: a method is just a function with a receiver argument.
Here's Abs written as a regular function with no change in functionality.
*/
func GoMethodsAreFunctions() {
	fmt.Printf("****Running methods.GoMethodsAreFunctions(), go methods are nothing but function with receiver argument \n")
	v := vertex3{3, 4}
	fmt.Println(abs(v))
}

type myFloat float64

func (f myFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

/*
MethodsOnNonStructTypes - You can declare a method on non-struct types, too.

In this example we see a numeric type MyFloat with an Abs method.
You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
*/
func MethodsOnNonStructTypes() {
	fmt.Printf("****Running methods.MethodsOnNonStructTypes(), You can declare a method on non-struct types, too.  \n")
	f := myFloat(-math.Sqrt2)
	fmt.Println(f.abs())
}

func (v vertex3) abs2() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *vertex3) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/*
MethodsWithPointerReceiver - You can declare methods with pointer receiver.

This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)

For example, the Scale method here is defined on *Vertex.

Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often
need to modify their receiver, pointer receivers are more common than value receivers.

Try removing the * from the declaration of the Scale function on line 16 and observe how the program's behavior changes.

With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for any
other function argument.) The Scale method must have a pointer receiver to change the Vertex value declared in the main function.
*/
func MethodsWithPointerReceiver() {
	fmt.Printf("****Running methods.MethodsWithPointerReceiver(), You can declare methods with pointer receiver.  \n")
	v := vertex3{3, 4}
	v.scale(10)
	fmt.Println(v.abs2())
}
