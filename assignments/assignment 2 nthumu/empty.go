package main

import (
	"fmt"
)

func main() {
	var bb [2]interface{}
	bb[0] = "hello"
	bb[1] = 9
	fmt.Println(bb)
}
