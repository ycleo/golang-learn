//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

//* Using an array, create a shopping list with enough room for 4 products
//  - Products must include the price and the name
type Product struct {
	name  string
	price float32
}

//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

func printListInfo(list [4]Product) {
	var totalItem int
	var totalCost float32

	for i := 0; i < len(list); i++ {
		item := list[i]
		if item.name != "" {
			fmt.Println(item.name, item.price)
			totalItem += 1
			totalCost += item.price
		}
	}
	fmt.Println("Last item", list[totalItem-1])
	fmt.Println("Total items:", totalItem)
	fmt.Println("Total cost:", totalCost)
}

func main() {
	shoppingList := [4]Product{
		{"Apple", 5.6},
		{"Orange", 7.1},
		{"Banana", 6.6},
	}

	printListInfo(shoppingList)

	shoppingList[3].name = "Grape"
	shoppingList[3].price = 3.2

	printListInfo(shoppingList)
}
