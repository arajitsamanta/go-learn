package flowcontrol

import "fmt"

/*
SimpleForLoop - Go simple for loop example.

Go has only one looping construct, the for loop.
The basic for loop has three components separated by semicolons:
the init statement: executed before the first iteration
the condition expression: evaluated before every iteration
the post statement: executed at the end of every iteration
The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.
The loop will stop iterating once the boolean condition evaluates to false.
Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.
*/
func SimpleForLoop() {
	fmt.Printf("\n****Running flowcontrol.SimpleForLoop(), go for loop example ")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//ForOptionalInitAndPost - Go for loop with out init and post statement
func ForOptionalInitAndPost() {
	fmt.Printf("\n****Running flowcontrol.ForOptionalInitAndPost(), go for loop with out init and post statement example ")
	sum := 1
	for sum < 1000 {
		sum = sum + sum
	}
	fmt.Println(sum)
}

//ForeverLoop -  If you omit the loop condition it loops forever, so an infinite loop is compactly expressed
func ForeverLoop() {
	fmt.Printf("\n****Running flowcontrol.ForForever(), go infinte for loop")
	for {
	}
}
