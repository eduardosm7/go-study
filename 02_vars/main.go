package main

import "fmt"

func main() {

	var age int32 = 20
	const isCool = true

	name := "Eduardo"
	var size float32 = 1.3

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", size)
}
