//--Summary:
//  Create a program to manage lending of library books.
//

//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Title string // book title
type Name string  // library member name
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int // total owned by the library
	lended int // total books lended out
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}

		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
	}
}

func printMemberAudits(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ lended:", book.lended)
	}
	fmt.Println()
}

func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}
	if book.lended == book.total {
		fmt.Println("No more books available to lend")
		return false
	}
	book.lended += 1
	library.books[title] = book

	member.books[title] = LendAudit{checkOut: time.Now()}
	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of library")
		return false
	}

	audit, found := member.books[title]
	if !found {
		fmt.Println("Member did not check out this book")
		return false
	}

	book.lended -= 1
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}

//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time

func main() {
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	library.books["Romance"] = BookEntry{total: 5, lended: 0}
	library.books["Noval"] = BookEntry{total: 7, lended: 0}
	library.books["Poem"] = BookEntry{total: 10, lended: 0}
	library.books["History"] = BookEntry{total: 15, lended: 0}

	library.members["Shelly"] = Member{name: "Shelly", books: make(map[Title]LendAudit)}
	library.members["Ken"] = Member{name: "Ken", books: make(map[Title]LendAudit)}
	library.members["Jeff"] = Member{name: "Jeff", books: make(map[Title]LendAudit)}

	fmt.Println("Initial:")
	printLibraryBooks(&library)
	printMemberAudits(&library)
	fmt.Println("====================================================================")

	shelly := library.members["Shelly"]
	checkedOut := checkoutBook(&library, "History", &shelly)
	fmt.Println("\nCheck out a book:")
	if checkedOut {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}
	fmt.Println("====================================================================")

	returned := returnBook(&library, "History", &shelly)
	fmt.Println("\nReturn a book:")
	if returned {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}

}
