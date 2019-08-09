package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	//FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("FizzBuzz [%d]\n", i)
		} else if i%5 == 0 {
			fmt.Printf("Buzz [%d]\n", i)
		} else if i%3 == 0 {
			fmt.Printf("Fizz [%d]\n", i)
		} else {
			fmt.Println(i)
		}
	}
}
