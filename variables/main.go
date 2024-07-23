package main

import "fmt"

// another entry point, can execute this file
// you normally only have at most one func main per directory
// func main() {
// 	var myName string = "Max"
// 	fmt.Printf("Hello my name is %s\n", myName);
// }

func main () {
	myName := "Max" // inferred type, shorthand var decl
	myInt := 10
	myFloat := 10.0

	fmt.Printf("Hello my name is %s my int is %d my float is %f\n", myName, myInt, myFloat)

	// "zero value"
	var myFriendsName string // default ""
	var myBool bool // default false
	var myOtherInt int // default 0

	fmt.Printf("my other friends name %s my bool %t and my other int %d\n", myFriendsName, myBool, myOtherInt)

	myFriendsName = "Kyle"

	fmt.Printf("my other friends name %s my bool %t and my other int %d\n", myFriendsName, myBool, myOtherInt)
}