package main

import (
	"fmt"
	"encoding/json" 
)

type User struct {
	Name 	string `json:"name"`
	Address string `json:"address"`
}
	

func main() {

	var name string
	var address string

	fmt.Println("Start the make json program.")
	fmt.Printf("Please enter the name: ")
	fmt.Scanf("%s\n",&name)
	fmt.Printf("Please enter the address: ")
	fmt.Scanf("%s\n",&address)

	user := User{Name: name, Address: address}

	fmt.Printf("The name is %s\n", name)
	fmt.Printf("The address is %s\n", address)

	var j []byte
	j, err := json.MarshalIndent(user, "", " ")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(j))
	fmt.Println(user.Name)
	fmt.Println(user.Address)


	return
}
