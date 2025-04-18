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

	for i:=0; i<10; i++ {
		if i== 5  {
			continue
		}
		fmt.Print(i, "\n")
	}

	// SLICES AND ARRAYS, slices are dynamic arrays, so they can grow and shrink in size, while arrays are fixed in size.
	// Slices are more flexible and easier to work with than arrays, so they are used more often in Go. But slices work with arrays in the background, so they are not completely different.
	// Slices are a reference type, so when you pass a slice to a function, you are passing a reference to the underlying array, not a copy of the array. This means that if you modify the slice, you are modifying the underlying array as well. so if it changes in the funcution, it will change in the main function too.
	
	var array = [5]int{1,2,3,6,1}
	slice := array[1:3]
	fmt.Println(cap(array), len(array))
	fmt.Println(cap(slice), len(slice))
	slice = append(slice, 3, 1, 2)
	fmt.Println(cap(slice), len(slice))

	sliceN := []int{1,2,3,4,5}
	fmt.Println(cap(sliceN), len(sliceN))
	sliceN = append(sliceN, 3, 1, 2)
	fmt.Println(cap(sliceN), len(sliceN))

	sliceMade := make([]int, 5, 25)
	fmt.Println(cap(sliceMade), len(sliceMade))
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
