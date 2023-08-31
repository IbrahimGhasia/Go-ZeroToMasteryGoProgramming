package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello", name)
}

func message() string {
	return "Hello World! How are you?"
}

func sum(a,b,c int) int {
	return a+b+c
}

func main() {
	greet("Ibrahim")
	fmt.Println(message())
	fmt.Println(sum(1,2,3))
}