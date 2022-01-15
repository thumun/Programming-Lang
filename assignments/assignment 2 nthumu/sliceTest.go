package main

import "fmt"

type LS struct {
	dd int
}

func main() {
	a := make([]int, 3) // make a slice
	b := a
	a[0] = 5
	b[1] = 10
	fmt.Printf("a=%v\nb=%v\n", a, b)

	c := 5
	d := c
	fmt.Printf("c=%d   d=%d\n", c, d)
	c = 10
	d = 20
	fmt.Printf("c=%d   d=%d\n", c, d)

	ll := LS{5}
	llc := ll
	fmt.Printf("BEFORE  ll:%+v  llc:%+v\n", ll, llc)
	ll.dd = 7
	fmt.Printf("AFTER   ll:%+v  llc:%+v\n", ll, llc)
}