package main

import (
	"fmt"
	"strings"
)


func main() {
	var input string
	fmt.Scan(&input);
	if input[0] != 'i' {
		fmt.Println("Not Found!")
		return
	} 
	if input[len(input) - 1] != 'n' {
		fmt.Println("Not Found!")
		return
	}
	if !strings.Contains(input, "a") {
		fmt.Println("Not Found!")
		return
	}
	fmt.Println("Found!")
}
