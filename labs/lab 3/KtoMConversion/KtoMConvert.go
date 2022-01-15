package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	input := os.Args

	if input[2] == "k" {
		output, _ := strconv.ParseFloat(input[1], 64)
		output = output / 1.62137119
		fmt.Printf("%.2f miles", output)

	} else if input[2] == "m" {
		output, _ := strconv.ParseFloat(input[1], 64)
		output = output * 1.62137119
		fmt.Printf("%.2f k", output)

	} else {
		fmt.Printf("unknown input")
	}
}
