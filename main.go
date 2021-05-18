package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// variable that stores current opt
var opt string // **

// declare reader as a pointer from bufio
var reader *bufio.Reader

type User struct {
	id       int
	username string
	email    string
	age      int
}

var id int = 0
var users map[int]User

func createUser() {

	clearLine()

	// read values for creating user

	fmt.Print("Insert value for username: ")

	username := readLine(opt)

	fmt.Print("Insert value for email: ")
	email := readLine(opt)

	fmt.Print("Insert value for age: ")

	// convert string to int and assign err if ocurrs
	age, err := strconv.Atoi(readLine(opt))

	if err != nil {
		// throw the error
		panic("[Error] Unable to convert from string to int")
	}

	user := User{username: username, email: email, age: age}
	id++
	users[id] = user

	fmt.Println("[+] User data: ", users[id])
	fmt.Println("[+] User created succesfully")
}

func readUser() {
	clearLine()

	fmt.Println("[+] Current User list:")

	// loops into every user in range and prints that user
	fmt.Println("[-] ***Id***|***Username***|***Email***|***Age***")
	for id, user := range users {
		fmt.Println(
			"[*]",
			id, " - ",
			user.username,
			"/",
			user.email,
			" : ",
			user.age)
	}
}

func updateUser() {
	fmt.Print("Input the user id to update: ")
	id, err := strconv.Atoi(readLine(opt))

	if err != nil {
		panic("[Error] Unable to convert from string to int")
	}

	if _, ok := users[id]; ok {

		fmt.Print("Insert value for username: ")

		username := readLine(opt)

		fmt.Print("Insert value for email: ")
		email := readLine(opt)

		fmt.Print("Insert value for age: ")

		// convert string to int and assign err if ocurrs
		age, err := strconv.Atoi(readLine(opt))

		if err != nil {
			// throw the error
			panic("[Error] Unable to convert from string to int")
		}

		user := users[id]

		user.username = username
		user.email = email
		user.age = age

		users[id] = user

		fmt.Println("[+] User was updated succesfully")
	}
}

func deleteUser() {
	clearLine()
	fmt.Print("Input the user id to delete: ")
	id, err := strconv.Atoi(readLine(opt))

	if err != nil {
		panic("[Error] Unable to convert from string to int")
	}

	// if user by id is found
	if _, ok := users[id]; ok {
		// delete user
		delete(users, id)

	}
	fmt.Println("[+] Deleted user succesfully")
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

/**

	Clear console

**/
func clearLine() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	clearLine()
	users = make(map[int]User)

	// Define a buffer to read from standard input
	reader = bufio.NewReader(os.Stdin)
	for {
		if opt == "quit" || opt == "q" || opt == "^C" {
			break
		}
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
	fmt.Println("Closing...")
}
