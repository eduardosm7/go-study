package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func square(num1 int) int {
	return num1 * num1
}

func main() {
	fmt.Println(greeting("Eduardo"))
	fmt.Println(getSum(3, 4))
	fmt.Println(square(5))
}
