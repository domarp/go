package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name string
	var address string
	var address_book map[string]string

	fmt.Printf("Enter name : ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		name = scanner.Text()
	}
	fmt.Println("Name entered = ", name)

	fmt.Printf("Enter address : ")
	if scanner.Scan() {
		address = scanner.Text()
	}
	fmt.Println("Address entered = ", address)

	address_book = map[string]string{
		"name":    name,
		"address": address,
	}

	barr, err := json.Marshal(address_book)
	if err == nil {
		fmt.Println("\nJSON : ", string(barr))
	} else {
		fmt.Println("Error : ", err)
	}

}
