//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//* Write a function that accepts a person's name as a function parameter and displays a greeting to that person.
func greet(name string) {
	fmt.Println("Hi", name)
}

//* Write a function that returns any message, and call it from within fmt.Println()
func showMessage(message string) string {
	return message
}

//* Write a function to add 3 numbers together, supplied as arguments, and return the answer
func addThreeNum(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

//* Write a function that returns any number
func five() int {
	return 5
}

//* Write a function that returns any two numbers
func twoSixes() (int, int) {
	return 6, 6
}

func main() {

	greet("Leo")

	fmt.Println(showMessage("What's up?"))

	//* Add three numbers together using any combination of the existing functions.
	//  * Print the result
	//* Call every function at least once

	firstSix, secondSix := twoSixes()
	sum := addThreeNum(five(), firstSix, secondSix)

	fmt.Println("The sum is", sum)
}
