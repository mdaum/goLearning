package main

import "fmt"

type Person struct {
	Name string
	Age int
}

// effectively constructor for Person, not required
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

	// lets actually pass ref of myPerson into the function that takes a reference, which will actually change the object
	changeAPersonName(&myPerson, "Cloud")
	fmt.Printf("this is my person %+v\n", myPerson)

	// lets use a receiver instead
	myPerson.changeName("Zack")
	fmt.Printf("this is my person %+v\n", myPerson)

}