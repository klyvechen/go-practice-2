package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"log"
)

type Student struct {
	fname string
	lname string
}

func main() {

	var filename string

	fmt.Println("Start give the file name:")
	fmt.Scanf("%s\n",&filename)
	
	f, err := os.Open(filename)

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)


	result := []Student{}

    for scanner.Scan() {
        name := scanner.Text()
		flname := strings.Split(name, " ")
		stu := Student{fname: flname[0], lname: flname[1]}
		result = append(result, stu)
    }

	for _, name := range result {
		fmt.Println(name)
	}

	return
}
