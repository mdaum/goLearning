package main

import (
	"fmt"
	"maxFemBasics/circle"
	"maxFemBasics/imports" //import syntax for Go: `${nameOfModule in go.mod}/${path}`
	// femBasicsImports "maxFemBasics/imports"  // example of aliasing an import. used to avoid collisions in the event something else is using "imports"
)

// always required
func main() {
	fmt.Println("hello world!")
	newTicket := imports.Ticket{
		ID:    123,
		Event: "FEM course",
	}
	newTicket.PrintEvent()

	myCircle := circle.NewCircle(5)
	fmt.Println(myCircle.ToString())
	fmt.Printf("myCircle Circumference is %f\n", myCircle.Circumference())
	fmt.Printf("myCircle Area is %f\n", myCircle.Area())
}
