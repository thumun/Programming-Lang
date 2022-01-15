package main

import (
	"fmt"
)

type Date struct {
	year  int
	month int
	day   int
}

func (adt Date) String() string {
	output := fmt.Sprintf("\tyear:%d month:%d day:%d \n",
		adt.year, adt.month, adt.day)
	return output
}

type Student struct {
	name     string
	id       int
	birthday Date
}

func main() {
	neha := Student{"neha", 0002, Date{2002, 04, 23}}
	devansh := Student{"devansh", 0003, Date{2001, 11, 27}}

	fmt.Printf("%v \n", neha)
	fmt.Printf("%v \n", devansh)
}
