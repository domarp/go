package main

import (
	"fmt"
)

func main() {
	var input_float float32
	fmt.Printf("Enter floating point number : ")
	fmt.Scan(&input_float)
	fmt.Println(int(input_float))
}
