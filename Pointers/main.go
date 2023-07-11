package main

import "fmt"

// Pointer(8 bytes integer for 64-bit computer, representing the address in the memory) is used in the following cases:
// - When we need to update the state/value of a variable.
// - When we want to optimize memory for large objects that are frequently called.
// However, using pointers everywhere can lead to errors. For example:
// If we have a pointer setting a user object across the application and update it somewhere, it will get updated everywhere.
// Debugging such scenarios can be challenging, and dealing with nil pointer references can be annoying and hard to debug.
// Pointers can also introduce inconsistencies, especially during runtime. Therefore, it's recommended to use pointers judiciously.

type User struct {
	email    string
	username string
	age      int
	file     []byte // Large object, potentially containing a file
}

// The following two functions are exactly the same but written in different formats.

// Using a regular function with a parameter
/*
func Email(u User) string {
	return u.email
}
*/

// Using a method with a receiver
func (u User) Email() string {
	return u.email
}

// Uncomment for the first time
// The below function does not use a pointer, so it cannot update the value of email outside its scope.
/*
func (u User) updateEmail(email string) {
	u.email = email
}
*/

// By using a pointer receiver, the function can update the email field of the User outside its scope.
// When using a pointer, only 8 bytes (size of a pointer) are copied into the function, regardless of the size of the User object.
// This optimization is beneficial when the User object contains large data, such as a file.
func (u *User) updateEmail(email string) {
	u.email = email
}

// This function takes a User object as a parameter and returns its email field.
// It does not modify the User object, so using a pointer receiver is not necessary.
func Email(user User) string {
	return user.email
}

func main() {
	user := User{
		email: "sumit@mail.com",
	}
	user.updateEmail("amit@mail.com")
	// fmt.Println("When we don't use pointers, it will print the same email: ", user.Email()) // Uncomment for the first time
	fmt.Println("When we use pointers, it will print the updated email: ", user.Email())
}
