// ex14
package main

import (
	"fmt"
)

func main() {
	fmt.Println("let's see if these two identical strings evaluate to equal")
	s1 := "hi"
	s2 := "hi"
	// okay, I showed a little conditional flow here
	if s1 == s2 {
		fmt.Printf("yes, they are!")
	}
}
