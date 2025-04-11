package main

import "fmt"
import "rsc.io/quote"

func main() {
    fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	
	var pisello string = "pisello"

	fmt.Println(pisello)
	fmt.Println(Hello(pisello))
}

func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}