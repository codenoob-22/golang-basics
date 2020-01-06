package main

import "fmt"

func main() {
	//arrays declarations

	fruit := [2]string{"apple", "orange"}

	fruite := []string{"ji", "ki", "mi"}

	fmt.Println(fruit)
	fmt.Println(fruite)
	
	var ar [10]int  // arrays have fixed sizes and this is where slices come into play
	var arrslice []int   // this is a slice, these are built on some array, this can have variable length  functions related to it are  
	
	// Creating an array of size 10, slices it till index 5, and returns the slice reference this way you can use slices to
	//get some flexibility  when you want it to be like 
	s := make([]int, 5, 10)
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s)) // cap(s) tells us the capacity
	
	
	// you can make another slice to append value to previous array
	
	slice1 := []string{"C", "C++", "Java"}
	slice2 := append(slice1, "Python", "Ruby", "Go")

	fmt.Printf("slice1 = %v, len = %d, cap = %d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 = %v, len = %d, cap = %d\n", slice2, len(slice2), cap(slice2))

	slice1[0] = "C#"
	fmt.Println("\nslice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
	
	//array traversal methods
	
	//first  way  of iterating through  an array 
	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}

	// second way of iterating through array
	for index, element := range intArray {
		fmt.Println(index, "=>", element)

	}

	//third way of iterating through array
	for _, value := range intArray {
		fmt.Println(value)
	}

	j := 0
	//fourth way of iterating through array
	for range intArray {
		fmt.Println(intArray[j])
		j++
	}
}
