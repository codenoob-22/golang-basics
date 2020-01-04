package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 5, 6, 7, 8}

	//loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d \n", i, id)
	}

	//Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d \n", id)
	}

	//add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("sum:", sum)

	//range with map
	emails := map[string]string{"Bob": "bob", "bon": "bon", "nob": "nob"}

	for k, v := range emails {
		fmt.Printf("%s: %s \n", k, v)
	}
}
