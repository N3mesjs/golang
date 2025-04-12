package main

import (
	"fmt"
)

// Defining a custom error type because error is an interface
// type divideError struct {
// 	Divider int
// }
// func (de divideError) Error() string {
// 	return fmt.Sprintf("cannot divide by %d", de.Divider)
// }

func main() {
	res, err := divisions(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res)
	}
}

func divisions(a, b int) (int, error) {
	if a == 0 {
		return 0, fmt.Errorf("division by zero")
	} else if b == 0 {
		return 0, fmt.Errorf("division by zero")
	} else {
		result := a / b
		return result, nil
	}
}