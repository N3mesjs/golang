package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	
	for {
		fmt.Printf("Ore: %.0f, Minuti: %.0f, Secondi: %.0f\n", time.Since(start).Hours(), time.Since(start).Minutes(), time.Since(start).Seconds())
		time.Sleep(1 * time.Second)
	}
}