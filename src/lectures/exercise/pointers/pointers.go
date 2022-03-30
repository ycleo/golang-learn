//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:

package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
type Item struct {
	name string
	tag  SecurityTag
}

//* Create functions to activate and deactivate security tags using pointers
func activate(tag *SecurityTag) {
	*tag = Active
}
func deactivate(tag *SecurityTag) {
	*tag = Inactive
}

//* Create a checkout() function which can deactivate all tags in a slice
func checkout(itemList []Item) {

	for _, item := range itemList {
		deactivate(&item.tag)
	}
}

func main() {
	//* Perform the following:
	//  - Create at least 4 items, all with active security tags
	//  - Store them in a slice or array
	itemList := []Item{
		{"GDB", Active},
		{"Mouse", Active},
		{"Network", Active},
		{"OS", Active},
	}
	fmt.Println(itemList)
	//  - Deactivate any one security tag in the array/slice
	deactivate(&itemList[1].tag)
	fmt.Println(itemList)
	//  - Call the checkout() function to deactivate all tags
	checkout(itemList)
	fmt.Println(itemList)
	//  - Print out the array/slice after each change
}
