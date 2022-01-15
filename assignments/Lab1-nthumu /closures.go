package main

import (
	"fmt"
)


func main() {
	b6 := a(6)
	b42 := a(42)

	fmt.Printf("b6 :%v\n", b6(5))
	fmt.Printf("b42:%v\n", b42(7))
	fmt.Printf("b6 :%v\n", b6(5))
	fmt.Printf("b42:%v\n", b42(7))
	fmt.Printf("b6 :%v\n", b6(25))
	fmt.Printf("b42:%v\n", b42(-11))
}

// Not strictly necessary, but defining types for funtions being passed, or
// returned make the code more readable, IMHO
type af func(a int) int

/*
* Function that creates and returns a function.
* The created function simply returns the value passes into this
* function plus, the value passed into the created function.
* See the main method for examples
* @param int the value to increment
* @return the created function
 */
func a(aa int) af {
	b := func(incr int) int {
		return aa + incr
	}

	return b
}
