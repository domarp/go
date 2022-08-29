package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var integer_slice []int
	integer_slice = make([]int, 0, 3)

	for true {
		var input_string string
		fmt.Printf("Enter string : ")

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input_string = scanner.Text()
			input_integer, err := strconv.Atoi(input_string)

			if string(input_string) == "X" {
				fmt.Println("Exiting, final slice : ", integer_slice)
				break
			} else if err != nil {
				fmt.Println(input_string, "does not look like an integer, enter 'X' if you would like to exit")
				continue
			}
			integer_slice = append(integer_slice, input_integer)
			sort.Ints(integer_slice)
			fmt.Println(integer_slice)
		}
	}
}
