package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// declare reader as a pointer from bufio
var reader *bufio.Reader

func createUser() {
	fmt.Println("User created succesfully")
}

func readUser() {
	fmt.Println("User list:")
}

func updateUser() {
	fmt.Println("Updated user")
}

func deleteUser() {
	fmt.Println("Deleted user")
}

/**

	Function for reading input line

**/
func readLine(option string) string {
	// on error thrown panic
	if option, err := reader.ReadString('\n'); err != nil {
		panic("Cant find this option")
	} else {
		// Trim space from input and save to option
		return strings.TrimSpace(option)
	}
}

func main() {
	var opt string // **

	// Define a buffer to read from standard input
	reader = bufio.NewReader(os.Stdin)

	// Print options
	fmt.Println("A) Create")
	fmt.Println("B) Read")
	fmt.Println("C) Update")
	fmt.Println("D) Delete")

	// print and read input from line
	fmt.Print("Select a option to continue: ")

	// Pass option to readline and return value
	opt = readLine(opt)

	switch opt {
	case "a", "create":
		createUser()
	case "b", "read":
		readUser()
	case "c", "update":
		updateUser()
	case "d", "delete":
		deleteUser()
	default:
		fmt.Println("Invalid Option")
	}

}
