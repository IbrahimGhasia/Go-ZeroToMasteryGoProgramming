// Create a program to calculate area and perimeter of a rectangle

package main

import "fmt"

type Rectangle struct {
	width int
	height int
}

func area(r Rectangle) int {
	return r.width * r.height
}

func perimeter(r Rectangle) int {
	return 2 * (r.width + r.height)
}

func main() {
	rect := Rectangle{width: 5, height: 10}
	fmt.Println("Area:", area(rect))
	fmt.Println("Perimeter", perimeter((rect)))

	rect.width = 10
	rect.height = 20
	fmt.Println("Area:", area(rect))
	fmt.Println("Perimeter", perimeter((rect)))
}