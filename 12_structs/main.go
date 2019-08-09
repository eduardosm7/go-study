package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

func (person Person) greet() string {
	return "Hello, my name is " + person.firstName + " " + person.lastName + " and I am " + strconv.Itoa(person.age)
}

func (person *Person) hasBirthday() {
	person.age++
}

func (person *Person) getMarried(spouseLastName string) {
	if person.gender == "m" {
		return
	} else {
		person.lastName = spouseLastName
	}
}

func main() {

	person1 := Person{
		firstName: "Samantha",
		lastName:  "Smith",
		city:      "Boston",
		gender:    "f",
		age:       25,
	}

	person2 := Person{
		"Samantha",
		"Smith",
		"Boston",
		"f",
		25,
	}

	fmt.Println(person1)
	fmt.Println(person2)

	person1.hasBirthday()
	person1.getMarried("Williams")
	fmt.Println(person1.greet())

}
