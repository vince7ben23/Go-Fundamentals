package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	s := make([]int, 0)
	// fmt.Printf("%v, length: %v, capacity: %v \n", s, len(s), cap(s))
	var input string
	for {
		fmt.Printf("Enter a integer(enter \"X\" to exit): ")

		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if input == "X" {
			break
		}
		val, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		s = append(s, val)
		// fmt.Printf("capacity: %v\n", cap(s))
		sort.Ints(s)
		fmt.Println(s)
	}
}
