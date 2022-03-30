package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func ckeckCleanLiness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i] // good habbit for concurrency
		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is dirty")
		}
	}
}

func main() {
	rooms := [...]Room{
		{name: "Office"},
		{name: "Warehouse"},
		{name: "Reception"},
		{name: "Ops", cleaned: true},
	}
	ckeckCleanLiness(rooms)

	fmt.Println("cleaning the office...")
	rooms[0].cleaned = true

	ckeckCleanLiness(rooms)
}
