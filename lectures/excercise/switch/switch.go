package main

func main() {
	switch age := 2; {
	case age == 0:
		println("newborn")
	case age >=1 && age <= 3:
		println("toddler")
	case age < 13:
		println("child")
	case age < 18:
		println("teenager")
	default:
		println("adult")
	}
}