package types

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

/*
FunctionAsValues - Function are values too.

They can be passed around just like other values.Function values may be used as function arguments and return values
*/
func FunctionAsValues() {
	fmt.Printf("****Running types.FunctionAsValues(), running function as values \n")
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

/*
FunctionClosures - Go functions may be closures.

A closure is a function value that references variables from outside its body. The function may access
and assign to the referenced variables; in this sense the function is "bound" to the variables.

For example, the adder function returns a closure. Each closure is bound to its own sum variable.
*/
func FunctionClosures() {
	fmt.Printf("****Running types.FunctionClosures(), functions can have closures \n")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

}

func fibonacci() func() int {
	x := 0
	y := 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

/*
FibonacciWithFunctionClosure - fibonacci implementation with closure.

Let's have some fun with functions.
Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
*/
func FibonacciWithFunctionClosure() {
	fmt.Printf("****Running types.FibonacciWithFunctionClosure(), running fibonacci series with function closures \n")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
