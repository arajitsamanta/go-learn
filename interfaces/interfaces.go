package interfaces

import (
	"fmt"
	"math"
)

type abser interface {
	Abs() float64
}

type myFloat float64

func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type vertex struct {
	X, Y float64
}

func (v *vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
HelloInterfaces - An interface type is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

Note: There is an error in the example code on line 47. Vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex (the pointer type).
*/
func HelloInterfaces() {
	fmt.Printf("\n****Running interfaces.HelloInterfaces(), go interfaces \n")
	var a abser
	f := myFloat(-math.Sqrt2)
	v := vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v

	fmt.Println(a.Abs())
}

type interface1 interface {
	M()
}

type structT struct {
	S string
}

// M - This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t structT) M() {
	fmt.Println(t.S)
}

/*
ImplicitInterfaceImplementations - A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.
*/
func ImplicitInterfaceImplementations() {
	fmt.Printf("****Running interfaces.ImpliciteInterfaceImplementations(), go interface implemented implicitely, no \"implements\" keyword \n")
	var i interface1 = structT{"Go interfaces implemented implicitly"}
	i.M()
}
