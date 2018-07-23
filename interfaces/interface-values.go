package interfaces

import (
	"fmt"
	"math"
)

type interface2 interface {
	M()
}

type structT2 struct {
	S string
}

func (t *structT2) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type myF float64

func (f myF) M() {
	fmt.Println(f)
}

func describe(i interface2) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
InterfaceValues -Interface values.

Under the covers, interface values can be thought of as a tuple of a value and a concrete type:

(value, type)

An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface value executes the method of the same name on its underlying type.
*/
func InterfaceValues() {
	fmt.Printf("****Running interfaces.InterfaceValues(), Under the covers, interface values can be thought of as a tuple of a value and a concrete type \n")
	var i interface2

	i = &structT2{"Hello"}
	describe(i)
	i.M()

	i = myF(math.Pi)
	describe(i)
	i.M()
}

/*
InterfaceWithNilValues - If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M
in this example.)

Note that an interface value that holds a nil concrete value is itself non-nil.
*/
func InterfaceWithNilValues() {
	fmt.Printf("****Running interfaces.InterfaceWithNilValues(), If the concrete value inside the interface itself is nil, the method will be called with a nil receiver. \n")
	var i interface2

	var t *structT2
	i = t
	describe(i)
	i.M()

	i = &structT2{"hello"}
	describe(i)
	i.M()
}

/*
NilInterfaceValues - A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
*/
func NilInterfaceValues() {
	fmt.Printf("****Running interfaces.NilInterfaceValues(), A nil interface value holds neither value nor concrete type. \n")
	var i interface2
	describe(i)
	i.M()
}

func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*
EmptyInterfaces - The interface type that specifies zero methods is known as the empty interface:

	interface{}

An empty interface may hold values of any type. (Every type implements at least zero methods.)
Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.
*/
func EmptyInterfaces() {
	fmt.Printf("****Running interfaces.EmptyInterfaces(), interface with zero methods called Empty interface \n")
	var i interface{}
	describe2(i)

	i = 42
	describe2(i)

	i = "hello"
	describe2(i)
}
