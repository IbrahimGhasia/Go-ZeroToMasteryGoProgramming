package main

import "fmt"

func main() {
	for i:= 1; i <= 50; i++ {
		divisibleByThree := i % 3 == 0
		divisibleByFive := i % 5 == 0

		if divisibleByThree && divisibleByFive {
			fmt.Println("FizzBuzz")
		} else if divisibleByThree {
			fmt.Println("Fizz")
		} else if divisibleByFive {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}