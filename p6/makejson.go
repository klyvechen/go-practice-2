package main

import (
	"fmt"
	"encoding/json" 
)


func main() {

	var name string
	var address string

	fmt.Println("Start the make json program.")
	fmt.Printf("Please enter the name: ")
	fmt.Scanf("%s\n",&name)
	fmt.Printf("Please enter the address: ")
	fmt.Scanf("%s\n",&address)

	m := make(map[string]string)
	m["name"] = name
	m["address"] = address

	fmt.Printf("The name is %s\n", name)
	fmt.Printf("The address is %s\n", address)

	var j []byte
	j, err := json.MarshalIndent(m, "", " ")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(j))

	return
}
