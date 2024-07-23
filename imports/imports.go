package imports

import "fmt"

type Ticket struct { // note that all fields must be capitalized for it to be used as import!!
	ID int
	Event string 
}

func (t Ticket) PrintEvent() {
	fmt.Println(t.Event)
}