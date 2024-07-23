package main

import (
	"fmt"
	"slices"
)

func main() {
	// statically sized arrays
	// animals := [2]string{} // string array of size 2, you must allocate the size

	// animals[0] = "dog"
	// animals[1] = "cat"
	// // animals[2] = "ferret" // out of bounds

	// fmt.Println(animals);

	// slices, which are dynamically sized arrays
	animals := []string{
		"dog",
		"cat",
	} // slice of type string

	animals = append(animals, "moose")

	fmt.Println(animals)
	
	// to remove an elem from a slice you append two portions of the slice together, omitting the desired elem we wish to remove
	// or use the helper slices
	animals = slices.Delete(animals, 0, 1) // delete dog 

	fmt.Println(animals)

	animals = append(animals, "dog")

	// here is how you would delete cat w/o Delete method

	animals = append(animals[:0], animals[1:]...) // so this is saying append the elem after index 0 with every elem after index 1, uses the 

	fmt.Println(animals)

	for i:=0; i < len(animals); i++ {
		fmt.Printf("this is my animal %s\n", animals[i])
	}

	for i:=0; i < 10; i++{
		fmt.Println(i)
	}

	// can also use range as of go 1.22
	for index, value := range animals {
		fmt.Printf("this is my index %d and this is my animal %s\n", index, value)
	}

	for value := range 10 {
		fmt.Println(value)
	}

	// while loops?
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}