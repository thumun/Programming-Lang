/* Name:    Neha Thumu
 * Desc:
 *    The main driver program for Part 2 of Assignment 2.
 *    This program is seeing how stack depth changes based on the size of an
 * 	  array defined inside a recursive function.
 * Modified: Sep 28, 2021
 */

package main

import (
	"fmt"
)

// recursively adding an int - to test stack overflow (function version 1)
// func a(n int) int {
// 	if n < 1 {
// 		return n
// 	} else {
// 		fmt.Printf("\n%v\n", n)
// 		return a(n + 1)
// 	}
// }

var counter int64 = 1

/**
* Testing stack overflow by changing the size of the array and recursively
* adding a float
*
**/
func rec_test(n float64) float64 {
	var arr [200000]float64
	if n < 1 {
		return n
	} else {
		fmt.Printf("\n%d\n", counter)
		counter = counter + 1
		return rec_test(n + arr[0])
	}
}

/**
* Prof Towell's code for pt.2
**/
func r(ii int) {
	var ij [50000]int64

	if ii%100 == 0 {
		fmt.Printf("D %d  %d\n", ii, ij[0])
	}
	r(ii + 1)
}

func main() {
	// a(5)
	rec_test(1.0)
	// r(5)
}

/**
On my computer, there is 1 gigabyte of stack space.

The value is consistent as the amount of memory used by the recursive
function.
When the array size is increased by factors of 2, the stack depth decreases by
approximately factors of 2.

My code:

50,000 array of float64 gave 1341 stack depth
100,000 array of float64 gave 671 stack depth
200,000 array of float64 gave 335 stack depth

Using Prof. Towell's code:

50,000 array of int64 gave ~130,000 stack depth
100,000 array of int64 gave ~60,000 stack depth
200,000 array of int64 gave ~30,000 stack depth


By increasing size of the array reduces the stack depth in Go but this would
not lead to the same outcome in Java since Java uses the heap to allocate
space in memory.

**/
