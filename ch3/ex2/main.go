// ex. 2 else
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// use the rand library to generate a random integer.
	n := rand.Intn(100)
	// 'if' condition
	if n%2 == 0 {
		// code block that's executed if the condition is true
		fmt.Printf("Number is even: %d", n)
	} else {
		// code block that's executed if it isn't true
		fmt.Printf("Number is odd: %d", n)
	}

}
