package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input_string string
	fmt.Printf("Enter string : ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input_string = scanner.Text()
	}

	lower_string := strings.ToLower(input_string)

	if (strings.Contains(lower_string, "a")) &&
		(strings.HasPrefix(lower_string, "i")) &&
		(string(lower_string[len(lower_string)-1]) == "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
