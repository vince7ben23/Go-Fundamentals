package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a string: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	// input = strings.TrimSuffix(input, "\r\n")
	input = strings.ToLower(input)

	if input[0:1] == "i" && input[len(input)-1:] == "n" && strings.Contains(input, "a") {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}

}
