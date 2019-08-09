package main

import "fmt"

func main() {
	emails := make(map[string]string)

	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	delete(emails, "Bob")

	fmt.Println(emails)

	emails2 := map[string]string{
		"Bob":    "bob@gmail.com",
		"Sharon": "sharon@gmail.com",
	}

	fmt.Println(emails2)

}
