package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	var filename string
	var slice_names []name

	max_char_size := 20

	// Read filename
	fmt.Printf("Enter filename : ")
	scanner1 := bufio.NewScanner(os.Stdin)
	if scanner1.Scan() {
		filename = scanner1.Text()
	}
	fmt.Println("File entered = ", filename)

	// Open file and read file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Scan each line & append to slice of structs
	scanner2 := bufio.NewScanner(file)
	for scanner2.Scan() {
		fullname := strings.Fields(scanner2.Text())
		trimmed_fname := fullname[0]
		trimmed_lname := fullname[1]

		if len(fullname[0]) > max_char_size {
			trimmed_fname = fullname[0][:20]
		}
		if len(fullname[1]) > max_char_size {
			trimmed_lname = fullname[1][:20]
		}

		slice_names = append(slice_names, name{trimmed_fname, trimmed_lname})

	}
	if err := scanner2.Err(); err != nil {
		log.Fatal(err)
	}

	// Iterate slice of structs and print first and last name
	for _, names := range slice_names {
		fmt.Println("\n\nFirst Name : ", names.fname, "\nLast Name : ", names.lname)
	}
}
