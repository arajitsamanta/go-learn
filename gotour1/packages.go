package gotour1

/*
Every Go program is made up of packages.

Programs start running in package main.

This program is using the packages with import paths "fmt" and "math/rand".

By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.

Note: the environment in which these programs are executed is deterministic, so each time you run the example program rand.Intn will return the same number. (To see a different number, seed the number generator;
see rand.Seed. Time is constant in the playground, so you will need to use something else as the seed
*/
import (
	"fmt"
	"math/rand"
)

//TestPackages function
func TestPackages() {
	fmt.Println("\n****Running gotour1.TestPackages(), My favourite no is", rand.Intn(10))
}
