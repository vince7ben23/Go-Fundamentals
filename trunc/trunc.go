package main

import (
	"fmt"
	"math"
)

func main() {
	var input float64
	var res int
	fmt.Printf("Enter a floating point number: ")
	fmt.Scan(&input)
	if input >= 0 {
		res = int(math.Floor(input))
	} else {
		res = int(math.Ceil(input))
	}
	fmt.Printf("Truncated number: ")
	fmt.Print(res)
}
