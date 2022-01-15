/* Name:    Neha Thumu
 * Desc:
 *    The main driver program for Part 1 of Assignment 2.
 *    This program shows that Go uses static scoping through the use of
 *	  function returns and printing the value of local. This shows that the
 *	  local var remains unaltered even when the function attempts to change
*	   the value of it.
 * Modified: Sep 28, 2021
*/

package main

import (
	"fmt"
)

// global variable (structure found on golang)
var scopeTest int = 23

func a() int {
	/**
	* simply returns the value (global var) scopeTest
	**/
	return scopeTest
}

func b() int {
	/**
	* sets value of (gloval var) scopeTest as 5 then returns
	**/
	scopeTest = 5
	return scopeTest
}

func main() {
	var scopeTest int = 10
	fmt.Printf("printing scopeTest (local): %v \n", scopeTest)

	aOutput := a()
	fmt.Printf("printing scopeTest a (global): %v \n", aOutput)

	fmt.Printf("printing scopeTest (local): %v \n", scopeTest)

	bOutput := b()
	fmt.Printf("printing scopeTest b (global): %v \n", bOutput)

	fmt.Printf("printing scopeTest (local): %v \n", scopeTest)
}

/*
why and how they prove the form of scoping used by Go?

Even though there is the same variable name (scopeTest) for the local and
global variables. Since Go uses static scoping, the local variable value
does not change when the functions are called. (Function a returns scopeTest
which is the global variable while Function b changes the value of scopeTest.
As seen by the print statements after function b, the local variable remains
unchanged while the global changes to 5.)

If Go used dynamic scoping, the last print statement would have printed 5 since
function b would have changed it.
*/
