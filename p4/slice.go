package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Start the slice program.")
	sli := make([]int, 2, 2)
	for k := 0; k < 100; k++ {
		fmt.Printf("---------------------\n")
		fmt.Printf("Please add the number to the slice: ")
		var s string
		fmt.Scanf("%s\n",&s)
		fmt.Printf("The input is %s\n", s)
		if s == "x" {
			break
		}

		if _, err := strconv.ParseInt(s, 10, 64); err != nil{
			fmt.Printf("%s is not an int, please enter another value\n", s)
			continue
		}
		if i, err := strconv.Atoi(s); err == nil {
			sli = append(sli, i)
			sort.Ints(sli)
			fmt.Print("The slice is ")
			fmt.Println(sli)
			fmt.Printf("The cap is %d, len is %d\n", cap(sli), len(sli))
		}
	}
	return
}
