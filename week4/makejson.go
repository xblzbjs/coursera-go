/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/

package main

import (
	"encoding/json"
	"fmt"
)

// Person contain name and
type Person struct {
	Name    string
	Address string
}

func main() {
	var p1 = new(Person)
	fmt.Println("Please enter a name:")
	fmt.Scan(&p1.Name)
	fmt.Println("Please enter a address:")
	fmt.Scan(&p1.Address)
	personMap := map[string]string{"name": p1.Name, "address": p1.Address}
	// b, err := json.Marshal(p1)
	b, err := json.Marshal(personMap)
	if err != nil {
		panic(err)
	}
	fmt.Println("Print the JSON object:", string(b))
}
