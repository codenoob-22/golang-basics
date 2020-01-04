package main

import "fmt"

func main() {
	x := 5
	y := 6

	if x < y {
		fmt.Printf(" %d is less than %d", x, y)
	} else {
		fmt.Printf(" %d is less than %d", y, x)
	}
	color := "red"

	if color == "red" {
		fmt.Println(" red")
	} else if color == "blue" {
		fmt.Println("blue")
	} else {
		fmt.Println("none")
	}

	switch color {
	case "red":
		fmt.Println("red")
	case "blue":
		fmt.Println("blue")
	default:
		fmt.Println("none")
	}
}
