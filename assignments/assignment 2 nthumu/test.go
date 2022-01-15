// Program
package main

import (
	"fmt"
	"sort"
)

type lab4 struct {
	theInt int
	theFl  float64
}

func (obj lab4) String() string {
	output := fmt.Sprintf("My own ****: %v, %v", obj.theInt, obj.theFl)
	return output
}

type sliceIntSort []lab4

func (S sliceIntSort) Len() int           { return len(S) }
func (S sliceIntSort) Swap(i, j int)      { S[i], S[j] = S[j], S[i] }
func (S sliceIntSort) Less(i, j int) bool { return S[i].theInt < S[j].theInt }

type sliceFlSort sliceIntSort

func (S sliceFlSort) Len() int           { return len(S) }
func (S sliceFlSort) Swap(i, j int)      { S[i], S[j] = S[j], S[i] }
func (S sliceFlSort) Less(i, j int) bool { return S[i].theFl < S[j].theFl }

type funcInt func(int) int

func closureAdd(input int) (funcInt, funcInt) {

	adder := func(a int) int {
		return input + a
	}

	subtractor := func(b int) int {
		return input - b
	}

	return adder, subtractor
}

func main() {
	var inst1 lab4 = lab4{1, 2.5}
	var inst2 lab4 = lab4{2, 1.25}

	aa := make(sliceIntSort, 0)
	aa = append(aa, inst1, inst2)

	fmt.Printf("Before %v\n", aa)
	sort.Sort(sliceIntSort(aa))
	fmt.Printf("Int sort:%v\n", aa)

	bb := sliceFlSort(aa)

	sort.Sort(bb)
	fmt.Printf("Float sort:%v\n", bb)

	add5, sub5 := closureAdd(5)

	fmt.Printf("addNum val: %v, sunVal:%v", add5(1), sub5(1))

}
