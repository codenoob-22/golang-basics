package main

import "fmt"

func greeting(name string) string {
	return "Hello" + name
}

func getsum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(greeting("michael"))
}
