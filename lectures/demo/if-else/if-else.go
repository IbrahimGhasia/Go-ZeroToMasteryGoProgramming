package main

import "fmt"

func average(a,b,c int) float32 {
	return float32(a+b+c)/3
}

func main() {
	quiz1, quiz2, quiz3 := 9, 7, 8

	if quiz1 > quiz2 {
		fmt.Println("quiz1 is greater than quiz2")
	} else if quiz1 < quiz2 {
		fmt.Println("quiz1 is less than quiz2")
	} else {
		fmt.Println("quiz1 is equal to quiz2")
	}

	if average(quiz1, quiz2, quiz3)  > 7 {
		fmt.Println("You passed the exam")
	} else {
		fmt.Println("You failed the exam")
	}
}