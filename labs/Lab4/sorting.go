package main

import (
	"fmt"
	"sort"
)

type sortingStruct struct {
	a int
	b float64
	c string
}

type b []sortingStruct // changing int val

// p is variable (the slice), b is the type (the struct)
func (p b) Len() int {
	return len(p)
}

// assuming i and j are indices
func (p b) Swap(i, j int) {
	temp := p[i]
	p[i] = p[j]
	p[j] = temp
}

func (p b) Less(i, j int) bool {
	if p[i].a < p[j].a {
		return true
	} else {
		return false
	}
}

type bb []sortingStruct // changing float val

// p is variable (the slice), b is the type (the struct)
func (p bb) Len() int {
	return len(p)
}

// assuming i and j are indices
func (p bb) Swap(i, j int) {
	temp := p[i]
	p[i] = p[j]
	p[j] = temp
}

func (p bb) Less(i, j int) bool {
	if p[i].b < p[j].b {
		return true
	} else {
		return false
	}
}

type bbb []sortingStruct // changing string val

// p is variable (the slice), b is the type (the struct)
func (p bbb) Len() int {
	return len(p)
}

// assuming i and j are indices
func (p bbb) Swap(i, j int) {
	temp := p[i]
	p[i] = p[j]
	p[j] = temp
}

func (p bbb) Less(i, j int) bool {
	if p[i].c < p[j].c {
		return true
	} else {
		return false
	}
}

// copy the functions and change for the bb and bbb types
// doing this b/c faster for runtime (instead of using if cases to check type)

func main() {

	s := make(b, 10)

	s[0] = sortingStruct{9, 23.65, "dragon"}
	s[1] = sortingStruct{19, 56.21, "wyvern"}
	s[2] = sortingStruct{87, 3.864, "serpent"}
	s[3] = sortingStruct{43, 90.321, "lizard"}
	s[4] = sortingStruct{1, 45.271, "gecko"}
	s[5] = sortingStruct{22, 6.478, "salamander"}
	s[6] = sortingStruct{34, 48.934, "hydra"}
	s[7] = sortingStruct{52, 1.92, "snake"}
	s[8] = sortingStruct{61, 64.81, "leviathan"}
	s[9] = sortingStruct{92, 91.0, "kraken"}

	fmt.Printf("\n sorted int: %v", s)
	sort.Sort(s)
	fmt.Printf("\n unsorted int: %v", s)

	fmt.Printf("\n")

	changeFloat := bb(s)
	fmt.Printf("\n unsorted float: %v", changeFloat)
	sort.Sort(changeFloat)
	fmt.Printf("\n sorted float: %v", changeFloat)

	fmt.Printf("\n")

	changeStr := bbb(s)
	fmt.Printf("\n unsorted str: %v", changeStr)
	sort.Sort(changeStr)
	fmt.Printf("\n sorted str: %v", changeStr)
}
