package circle

import (
	"fmt"
	"math"
)

const PI float64 = 3.14

type Circle struct {
	radius float64
}

func NewCircle(radius float64) *Circle { // we will return reference to the allocated Circle
	var newCircle = Circle{
		radius: radius,
	}
	return &newCircle
}

func (c *Circle) ToString() string {
	return fmt.Sprintf("{radius: %f}", c.radius)
}

func (c *Circle) Circumference() float64 {
	return 2 * PI * c.radius
}

func (c *Circle) Area() float64 {
	return PI * math.Pow(c.radius, 2)
}
