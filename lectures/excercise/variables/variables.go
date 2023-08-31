package main

import "fmt"

func main() {
	var favouriteColour = "blue"
	var birthYear, age = 2002, 21

	fmt.Println("My Favourite Colour is",favouriteColour)
	fmt.Println("My birth year is",birthYear, "and my age is",age)

	var (
		firstName = "Ibrahim"
		lastName = "Ghasia"
	)

	fmt.Println("My first name is",firstName, "and my last name is",lastName)

	var ageInDays int
	ageInDays = age * 365
	fmt.Println("My age in days is",ageInDays)
}