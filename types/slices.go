package types

import (
	"fmt"
	"strings"
)

/*
HelloSlices -

An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.
The type []T is a slice with elements of type T.
A slice is formed by specifying two indices, a low and high bound, separated by a colon:

a[low : high]

This selects a half-open range which includes the first element, but excludes the last one.
The following expression creates a slice which includes elements 1 through 3 of a:

a[1:4]
*/
func HelloSlices() {
	fmt.Printf("****Running types.HelloSlices(), go slices \n")
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s = primes[1:4]
	fmt.Println(s)
}

/*
SlicePointers - A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.
Other slices that share the same underlying array will see those changes.
*/
func SlicePointers() {
	fmt.Printf("****Running types.SlicePointers(), slice are like pointers \n")
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

/*
SliceLiterals - A slice literal is like an array literal without the length.

This is an array literal:
[3]bool{true, true, false}

And this creates the same array as above, then builds a slice that references it:
[]bool{true, true, false}
*/
func SliceLiterals() {
	fmt.Printf("****Running types.SliceLiterals(), slice literals \n")
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

/*
SliceDefaults - When slicing, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the length of the slice for the high bound.
For the array
var a [10]int
these slice expressions are equivalent:
a[0:10]
a[:10]
a[0:]
a[:]
*/
func SliceDefaults() {
	fmt.Printf("****Running types.SliceDefaults(), slice defaults \n")
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

/*
SliceLengthAndCapacity - A slice has both a length and a capacity.

The length of a slice is the number of elements it contains.The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in
the example program to extend it beyond its capacity and see what happens.
*/
func SliceLengthAndCapacity() {
	fmt.Printf("****Running types.SliceLengthAndCapacity(), slice length and capacity \n")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	// SHould thorw error --> panic: runtime error: slice bounds out of range
	//s = s[:5]
	//printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

/*
NilSlices - The zero value of a slice is nil.
A nil slice has a length and capacity of 0 and has no underlying array.
*/
func NilSlices() {
	fmt.Printf("****Running types.NilSlices(), slice with zero length and capacity \n")
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

/*
SliceWithMake - Creating a slice with make

Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
The make function allocates a zeroed array and returns a slice that refers to that array:
a := make([]int, 5)  // len(a)=5
To specify a capacity, pass a third argument to make:
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4

*/
func SliceWithMake() {
	fmt.Printf("****Running types.SliceWithMake(), create slice with built-in make() function \n")
	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
}

//SlicesOfSlice - Slices can contain any type, including other slices.
func SlicesOfSlice() {
	fmt.Printf("****Running types.SlicesOfSlice(), slices of slice \n")
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

/*
SliceAppend - Appending to a slice

It is common to append new elements to a slice, and so Go provides a built-in append function. The https://golang.org/pkg/builtin/#append of the built-in package
describes append.
func append(s []T, vs ...T) []T

The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.
The resulting value of append is a slice containing all the elements of the original slice plus the provided values.

If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.
(To learn more about slices, read the https://blog.golang.org/go-slices-usage-and-internals)
*/
func SliceAppend() {
	fmt.Printf("****Running types.SliceAppend(), appending to slice using built-in append() \n")
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

/*
RangeOverSlice - The range form of the for loop iterates over a slice or map.
When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
*/
func RangeOverSlice() {

	fmt.Printf("****Running types.RangeOverSlice(), using range over slice \n")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	//You can skip the index or value by assigning to _.
	//If you only want the index, drop the ", value" entirely.
	var pow2 = make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // == 2**i
	}

	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}

}
