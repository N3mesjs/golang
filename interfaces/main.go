package main

import (
	"fmt"
	"math"
)

/*
	So here we have 2 structs, Rectangle and Circle, and we want to calculate the area of each shape.
	We can do this by creating an interface called Shape, which has a method called area. in this case, the method is a function that returns a float64 value.
	Then we implement the area method for each struct, later we can create a function that takes a Shape as an argument and calls the area method on it. and the result will change depending on the type of shape we pass to it.
*/

type Rectangle struct {
	Length int
	Width  int
}

type Circle struct {
	Radius int
}

type Shape interface {
	area() float64
}

func main(){
	rect := Rectangle{
		Length: 10,
		Width:  5,
	}
	circle := Circle{
		Radius: 7,
	}

	printArea(rect)
	printArea(circle)
}

func (r Rectangle) area() float64 {
	return float64(r.Length * r.Width)
}

func (c Circle) area() float64 {
	return math.Pi * float64(c.Radius^2)
}

func printArea(s Shape) {
	fmt.Printf("Area: %f\n", s.area())

	// Type assertion to check if s is a Rectangle, If it is, the ok is boolean value with true
	rec, ok := s.(Rectangle)
	if ok {
		fmt.Println("It's a rectangle with length:", rec.Length, "and width:", rec.Width)
	} else {
		fmt.Println("It's not a rectangle")
	}

	// or we can do it with a type switch statement, only in a switch stamentent we can use: interface.(type) so we can check the type of the interface

	switch v := s.(type) {
	case Rectangle:
		fmt.Println("It's a rectangle with length:", v.Length, "and width:", v.Width)
	case Circle:
		fmt.Println("It's a circle with radius:", v.Radius)
	default:
		fmt.Println("Unknown shape")
	}
}