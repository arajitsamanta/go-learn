package interfaces

import (
	"fmt"
	"time"
)

type myError struct {
	When time.Time
	What string
}

func (e *myError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &myError{
		time.Now(),
		"it didn't work",
	}
}

/*
GoErrors - Go programs express error state with error values.

The error type is a built-in interface similar to fmt.Stringer:

type error interface {
    Error() string
}

(As with fmt.Stringer, the fmt package looks for the error interface when printing values.)
Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.

i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}

fmt.Println("Converted integer:", i)
A nil error denotes success; a non-nil error denotes failure.

*/
func GoErrors() {
	fmt.Printf("****Running interfaces.GoErrors(), Go errors example. \n")
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type errNegativeSqrt float64

func (e errNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

//Sqrt - simple sqrt function
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errNegativeSqrt(f)
	}
	return 0, nil
}

/*
ErrorsExcercise - Copy your Sqrt function from the earlier exercise and modify it to return an error value.

Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

Create a new type

	type ErrNegativeSqrt float64

and make it an error by giving it a

	func (e ErrNegativeSqrt) Error() string

method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".

Note: A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop. You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?

Change your Sqrt function to return an ErrNegativeSqrt value when given a negative number.
*/
func ErrorsExcercise() {
	fmt.Printf("****Running interfaces.ErrorsExcercise(), Go errors excercise. \n")
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
