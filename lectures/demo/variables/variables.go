package main

import "fmt"

func main() { 
	var myName = "Ibrahim"
	fmt.Println("My name is",myName)

	var name string = "Kate"
	fmt.Println("name =", name)

	userName:="admin"
	fmt.Println("userName =", userName)

	var sum int
	fmt.Println("The sum is", sum)

	part1, other := 1,5
	fmt.Println("Part1 is", part1, "other is", other)

	part2, other := 2,0
	fmt.Println("Part2 is", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("The sum is", sum)

	var (
		lessonName = "variables"
		lessonType = "Demo"
	)

	fmt.Println("Lesson", lessonName, "is a", lessonType)

	word1, word2, _ := "Hello", "world", "!"
	fmt.Println(word1, word2)

}