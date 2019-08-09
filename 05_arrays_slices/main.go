package main

import "fmt"

func main() {

	var fruitArr [2]string

	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fruitArr2 := []string{"Apple", "Orange"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr2)
	fmt.Println(len(fruitArr2))
	fmt.Println(fruitArr[1:2])

}
