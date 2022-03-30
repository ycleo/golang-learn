//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

//* Create a rectangle structure containing its coordinates
type Coordinate struct {
	x, y int
}

type Rectangle struct {
	topRight, bottomLeft Coordinate
}

func length(rectangle Rectangle) int {
	return rectangle.topRight.x - rectangle.bottomLeft.x
}

func width(rectangle Rectangle) int {
	return rectangle.topRight.y - rectangle.bottomLeft.y
}

//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
func area(rectangle Rectangle) {
	area := length(rectangle) * width(rectangle)
	fmt.Println("the area is", area)
}

func perimeter(rectangle Rectangle) {
	perimeter := (length(rectangle) + width(rectangle)) * 2
	fmt.Println("the perimeter is", perimeter)
}

//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
func main() {
	rectangle := Rectangle{
		topRight:   Coordinate{3, 2},
		bottomLeft: Coordinate{0, 0},
	}
	area(rectangle)
	perimeter(rectangle)

	rectangle.topRight.y *= 2
	rectangle.topRight.x *= 2
	area(rectangle)
	perimeter(rectangle)
}
