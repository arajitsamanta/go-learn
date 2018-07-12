package types

import (
	"fmt"

	"golang.org/x/tour/pic"
)

//Pic - Custom function
func Pic(dx, dy int) [][]uint8 {
	// Allocate two-dimensioanl array.
	a := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		a[i] = make([]uint8, dx)
	}

	// Do something.
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			switch {
			case j%15 == 0:
				a[i][j] = 240
			case j%3 == 0:
				a[i][j] = 120
			case j%5 == 0:
				a[i][j] = 150
			default:
				a[i][j] = 100
			}
		}
	}
	return a
}

/*
SliceExcercise - Implement a function Pic.

func Pic(dx, dy int) [][]uint8{

}

It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your
picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
(You need to use a loop to allocate each []uint8 inside the [][]uint8.)
(Use uint8(intValue) to convert between types.)

*/
func SliceExcercise() {
	fmt.Printf("****Running types.SliceExcercise(), create slice with built-in make() function \n")
	pic.Show(Pic)
}
