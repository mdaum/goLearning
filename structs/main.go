package main

import "fmt"

type Person struct {
	Name string
	Age int
}

// effectively constructor for Person, not required
// why capitalized? Well this is actually making it "public" and importable elsewhere
func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age: age,
	}
}

// receiver, a method implemented on a type
// note that if p was of type Person (not reference), we'd be modifying a copy just like we do in changeACopyOfPersonName
func (p *Person) changeName(newName string) {
	p.Name = newName
}

// takes a copy of a person, will modify the copy
func changeACopyOfPersonName(person Person, newName string) {
	fmt.Println("address of copy ", &person.Name)
	person.Name = newName
}

// takes a reference to a person, not a copy of a person, so it will modify in place
func changeAPersonName(person *Person, newName string) {
	person.Name = newName
}



func main() {
	myPerson := NewPerson("Max", 30)

	fmt.Println(myPerson)

	// %v means any type, putting + in front prints the field name
	fmt.Printf("this is my person %+v\n", myPerson)

	fmt.Println("address of person.Name ", &myPerson.Name)
	changeACopyOfPersonName(myPerson, "Cloud")
	fmt.Printf("this is my person %+v\n", myPerson)	// it didn't change? thats becuase a COPY was passed to changeName, not a reference

	// lets actually pass pointer to myPerson into the function that takes a reference, which will actually change the object
	changeAPersonName(&myPerson, "Cloud")
	fmt.Printf("this is my person %+v\n", myPerson)

	// lets use a receiver instead
	myPerson.changeName("Zack")
	fmt.Printf("this is my person %+v\n", myPerson)

	x := 7
	y := &x // b is the memory address at which a is stored
	fmt.Printf("x is %d stored at %v\n", x, &x)
	*y = 9 // this is the same as saying x = 9, we are deferencing b, which is reference to x
	fmt.Printf("x is %d still stored at %v\n", x, y)

	// note that passing a copy into functions will copy everyting, and I mean everything. if you have a huge nested struct that takes up a lot of memory, it will be copied in its entirety if you pass by copy
	// depending on the size, you may want to pass by reference *even if* you have no intention of modifying the original to save memory
}