package main

import (
	"fmt"
	"strconv"
)

// define person struct
type Person struct {
	firstname string
	lastname  string
	city      string
	gender    string
	age       int
}

//gretting method

func (p Person) greet() string {
	return "hello, my name is " + p.firstname + " " + p.lastname + " I am " + strconv.Itoa(p.age)
}

// update age
func (p *Person) hasBday() {
	p.age++
}

//getsmarried (pointer reciever)
func (p *Person) getsmarried(spouseLastName string) {
	if p.gender == "f" {
		p.lastname = spouseLastName
	}
}
func main() {
	//init person
	person1 := Person{firstname: "sam", lastname: "smith", city: "bostn", gender: "f", age: 20}
	fmt.Println(person1.greet())
	person1.hasBday()
	fmt.Println(person1.greet())
	person1.getsmarried("mathew")
	fmt.Println(person1.greet())
}
