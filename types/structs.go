package types

import "fmt"

type vertex struct {
	X int
	Y int
}

//HelloStructures - A struct is a collection of fields
func HelloStructures() {
	fmt.Printf("****Running types.HelloStructures(), go structures \n")
	fmt.Println(vertex{1, 2})
}

// StructureFieldAccess - Struct fields are accessed using a dot.
func StructureFieldAccess() {
	fmt.Printf("****Running types.StructureFieldAccess(), structure field access \n")
	v := vertex{10, 20}
	v.X = 4
	fmt.Println(v.X)
}

/*
StructureFieldAccessUsingPointers - Struct fields can be accessed through a struct pointer.

To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome,
so the language permits us instead to write just p.X, without the explicit dereference.
*/
func StructureFieldAccessUsingPointers() {
	fmt.Printf("****Running types.StructureFieldAccessUsingPointers(), structure field access using pointers \n")
	v := vertex{11, 21}
	p := &v
	p.X = 1e9
	fmt.Println(v.X)
}

var (
	v1 = vertex{1, 2}  // has type Vertex
	v2 = vertex{X: 1}  // Y:0 is implicit
	v3 = vertex{}      // X:0 and Y:0
	p  = &vertex{1, 2} // has type *Vertex
)

/*
StructLiterals - A struct literal denotes a newly allocated struct value by listing the values of its fields.

You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)
The special prefix & returns a pointer to the struct value.
*/
func StructLiterals() {
	fmt.Printf("****Running types.StructLiterals(), structure literals \n")
	fmt.Println(v1, p, v2, v3)
}
