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

func abs4(v vertex3) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func scale(v *vertex3, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/*
MethodsWithPointerReceiver2 - You can declare methods with pointer receiver.

Here we see the Abs and Scale methods rewritten as functions.

Again, try removing the * from line 16. Can you see why the behavior changes? What else did you need to change for the example to compile?

*/
func MethodsWithPointerReceiver2() {
	fmt.Printf("****Running methods.MethodsWithPointerReceiver2(), methods with pointer receiver  \n")
	v := vertex3{3, 4}
	scale(&v, 10)
	fmt.Println(abs4(v))
}

func (v *vertex3) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func scaleFunc(v *vertex3, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/*
MethodsAndPointerIndirection - Example

Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:

var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK


while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK

For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically.
That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.
*/
func MethodsAndPointerIndirection() {
	fmt.Printf("****Running methods.MethodsAndPointerIndirection(), methods and pointer indirection  \n")
	v := vertex3{3, 4}
	v.Scale(2)
	scaleFunc(&v, 10)

	p := &vertex3{4, 3}
	p.Scale(3)
	scaleFunc(p, 8)

	fmt.Println(v, p)

}

func (v vertex3) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func absFunc(v vertex3) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
MethodsAndPointerIndirection2 - Example

The equivalent thing happens in the reverse direction.

Functions that take a value argument must take a value of that specific type:

var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!
while methods with value receivers take either a value or a pointer as the receiver when they are called:

var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
In this case, the method call p.Abs() is interpreted as (*p).Abs().
*/
func MethodsAndPointerIndirection2() {
	fmt.Printf("****Running methods.MethodsAndPointerIndirection2(), methods and pointer reverse indirection  \n")
	v := vertex3{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(absFunc(v))

	p := &vertex3{4, 3}
	fmt.Println(p.abs())
	fmt.Println(absFunc(*p))

}

func (v *vertex3) scale1(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *vertex3) abs1() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
MethodsWithValueOrPointerReceiver - There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

In this example, both Scale and Abs are with receiver type *Vertex, even though the Abs method needn't modify its receiver.

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)
*/
func MethodsWithValueOrPointerReceiver() {
	fmt.Printf("****Running methods.MethodsWithValueOrPointerReceiver(), chhose between value or pointer receiver \n")
	v := &vertex3{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.abs1())
	v.scale1(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.abs1())
}
