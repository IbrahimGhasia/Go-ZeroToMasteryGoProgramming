package main

import "fmt"

const (
	Monday = 1
	Tuesday = 2
	Wednesday = 3
	Thursday = 4
	Friday = 5
	Saturday = 6
	Sunday = 7
)

const (
	Admin = 10
	Manager = 20
	Contractor = 30
	Member = 40
	Guest = 50
)

func weekday(day int) bool {
	return day <= 4
}

func accessGranted() {
	fmt.Println("Access Granted")
}

func accessDenied() {
	fmt.Println("Access Denied")
}

func main() {
	today, role := Tuesday, Member

	if (role == Admin || role == Manager) {
		accessGranted()
	} else if(role == Contractor && !weekday(today)) {
		accessGranted()
	} else if (role == Member && weekday(today)) {
		accessGranted()
	} else if (role == Guest && (today == Monday || today == Wednesday || today == Friday)) {
		accessGranted()
	} else {
		accessDenied()
	}
}