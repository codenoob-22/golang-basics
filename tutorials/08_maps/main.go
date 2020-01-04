package main

import "fmt"

func main() {
	// Defining maps
	emails := map[string]string{"bhai": "bhai@gmail.com"}

	// assign keys
	emails["bob"] = "bob@gmail.com"
	emails["bog"] = "bob@gmail.com"
	emails["mike"] = "mike@gmail.com"
	fmt.Println(emails)

	// Delete from map
	delete(emails, "bob")
	fmt.Println(emails)
}
