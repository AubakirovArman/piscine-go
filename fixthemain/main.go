package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...\n")
	ptrDoor.state = "OPEN"
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...\n")
	ptrDoor.state = "CLOSE"
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?\n")
	if ptrDoor.state == "CLOSE" {
		return false
	} else {
		return true
	}
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?\n")
	if ptrDoor.state == "OPEN" {
		return false
	} else {
		return true
	}
}

type Door struct {
	state string
}

func main() {
	var door Door
	OpenDoor(&door)
	if IsDoorClose(&door) {
		OpenDoor(&door)
	}
	if IsDoorOpen(&door) {
		CloseDoor(&door)
	}
}
