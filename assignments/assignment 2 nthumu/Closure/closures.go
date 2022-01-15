/* Name:    Neha Thumu
 * Desc:
 *    The main driver program for Part 3 of Assignment 2.
 *    This program shows an unbounded iterator by using a closure.
 * Modified: Sep 28, 2021
 */

package main

import (
	"fmt"
	"os"
	"strconv"
)

/**
 * Type definition for a function (to use it as a closure)
**/
type pls func() int

/**
 * the incrementer function. Returns a function which multiplies the input by
 * a value, n.
 * This generates an unbounded series of numbers that are multiples of the
 * intial input.
**/
func makeIncrementer(input int) pls {

	var n int = 0
	f := func() int {
		o := input * n
		n += 1
		return o
	}

	return f
}

/**
 * test function which runs the makeIncrementer() infinitely
**/
func runForever() {
	inputThree := makeIncrementer(3)
	inputSeven := makeIncrementer(7)
	for {
		fmt.Println(inputThree())
		fmt.Println(inputSeven())
	}
}

/*
 * tested this function using a command line input to ensure proper values
 * were being outputted
 */
func runWithInput(maxPrint int) {
	inputThree := makeIncrementer(3)
	inputSeven := makeIncrementer(7)
	for i := 0; i < maxPrint; i++ {
		fmt.Println(inputThree())
		fmt.Println(inputSeven())
	}
}

func createIterator(iterator, iteratorTwo, maxPrint int) {
	input := makeIncrementer(iterator)
	for i := 0; i < maxPrint; i++ {
		fmt.Println(input())
	}
}

func main() {
	maxPrint, _ := strconv.Atoi(os.Args[1])
	runWithInput(maxPrint)

	n int := 8
	createIterator(n, m, maxPrint)

	//runForever()
}
