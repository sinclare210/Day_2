package main

import (
	"github.com/sinclare210/day_2/Book"
	"github.com/sinclare210/day_2/User"
)

// Define an interface that requires two methods: Display() and Save() error
// Any type (struct) that has these methods automatically satisfies this interface
type doer interface {
	Display()
	Save() error
}

// displayer takes anything that implements doer and calls its Display method
func displayer(data doer) {
	data.Display()
}

// saver takes anything that implements doer and calls its Save method
func saver(data doer) {
	data.Save()
}

func main() {
	// Create some book instances using the constructor function
	book1, _ := book.NewBook("Go in Action", "William Kennedy")
	book2, _ := book.NewBook("The Go Programming Language", "Alan Donovan")
	book3, _ := book.NewBook("Introducing Go", "Caleb Doxsey")

	// Create some user instances
	user1, _ := user.NewUser("Alice", "MIT", 21)
	user2, _ := user.NewUser("Bob", "Stanford", 22)
	user3, _ := user.NewUser("Charlie", "Harvard", 20)

	// Using the interface-based functions:
	// Both book1 and user1 can be passed into displayer() and saver()
	// because they implement the doer interface
	displayer(book1)
	displayer(user1)
	saver(book1)
	saver(user1)

	// Directly calling methods on other books and users
	user2.Display()
	user2.Save()
	book2.Display()
	book2.Save()

	user3.Display()
	user3.Save()
	book3.Display()
	book3.Save()
}
