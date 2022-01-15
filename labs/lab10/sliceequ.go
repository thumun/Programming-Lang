package main

import (
	"fmt"
	"reflect"
)

func main() {
	var aa = make([]interface{}, 10)
	for i := 0; i < 5; i++ {
		aa[i] = i
	}
	aa[5] = int32(32)
	aa[6] = "asdf"
	var bb = make([]interface{}, 5)
	for i := 0; i < 5; i++ {
		bb[i] = i
	}
	aa[7] = bb
	for i := 8; i < 10; i++ {
		aa[i] = int8('a' + i)
	}
	// fmt.Printf("%v\n", aa)
	// for _, v := range aa {
	// 	fmt.Printf("%v\n", sliceTest(v))
	// }

	var cc = make([]interface{}, 10)
	copy(cc, aa)

	output := sliceEquality(aa, cc)
	fmt.Printf("testing slices w/ same content: %v\n", output)

	cc[8] = "neha"
	output = sliceEquality(aa, cc)
	fmt.Printf("testing slices w/ diff content: %v\n", output)
}

// return true if the thing is a slice
func sliceTest(v interface{}) bool {
	rt := reflect.TypeOf(v)
	return rt.Kind() == reflect.Slice
}

// convert the thing passed to a slice
func sliceAssert(aa interface{}) []interface{} {
	return aa.([]interface{})
}

func sliceEquality(a []interface{}, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	} else {
		return recursionSlice(a, b)
	}
}

func recursionSlice(a []interface{}, b []interface{}) bool {
	if len(a) == 0 {
		return true
	} else {
		returnOut := compareSlice(a[0], b[0])

		if !returnOut {
			return false
		} else {
			return recursionSlice(a[1:], b[1:])
		}
	}
}

func compareSlice(valOne interface{}, valTwo interface{}) bool {
	var output bool = true

	if reflect.TypeOf(valOne) != reflect.TypeOf(valTwo) {
		output = false
	} else if sliceTest(valOne) && !sliceTest(valTwo) || !sliceTest(valOne) && sliceTest(valTwo) {
		output = false
	} else if sliceTest(valOne) && sliceTest(valTwo) {
		output = recursionSlice(sliceAssert(valOne), sliceAssert(valTwo))
	}

	return output
}
