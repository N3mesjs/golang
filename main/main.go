package main

import (
	"fmt"
)

func main() {
    //fmt.Println("Hello, World!")

	var i int32 = -32
	i *= 2

	greetings := "Hello, World!"
	fmt.Println("cc", greetings, i)

	fmt.Print(test(callback, 43, 54))

	// function that returns multiple values
	// in this case, the function returns two strings, but we only want the first name, so we use the blank identifier _ to ignore the second value, and it deletes the value from the memory
	firstName, _ := name()
	fmt.Println(firstName)

	x, y := coordinate()
	fmt.Println(x, y)
}

func callback(a, b int) int {
	return a+b
}

// function that take a function as an argument, the variable f has as type a function that takes two int arguments and returns an int
func test(f func(int,int) int, a, b int) int {
	return f(a,b)
}

func name() (string, string) {
	return "John", "Doe"
}

// i already declare the variables that the function will return, so i can only write the return statement without the variables
func coordinate() (x, y int){
	x = 1
	y = 2
	return
}

/* 
example of the normal way
func coordinate() (int, int) {
	x := 1
	y := 2
	return x, y
} 
*/
