package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits += 1 // if we have pointer to a struct, no need to deference
	fmt.Println("counter address (pointer value):", counter)
	fmt.Println("counter value:", *counter)
	fmt.Println("pointer address", &counter)
}
func replace(old *string, new string, counter *Counter) {
	*old = new

	fmt.Println("string address (pointer value):", old)
	fmt.Println("string value:", *old)
	fmt.Println("pointer address:", &old)
	fmt.Println()

	increment(counter)
}

func main() {
	counter := Counter{}
	hello := "Hello"
	world := "Wrold!"

	fmt.Println(hello, world)

	replace(&hello, "Hi", &counter)

	fmt.Println(hello, world)
	fmt.Println()

	phrase := []string{hello, world}
	fmt.Println(phrase)

	replace(&phrase[1], "Go!", &counter)
	fmt.Println(phrase)
}
