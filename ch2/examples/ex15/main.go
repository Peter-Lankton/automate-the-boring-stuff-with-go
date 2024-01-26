// ex15
package main

import (
	"fmt"
)

func main() {
	fmt.Println("let's see if these two non-identical strings evaluate to equal")
	s1 := "hi"
	s2 := "these are not the droids you're looking for"
	// okay, I showed a little conditional flow here
	if s1 != s2 {
		fmt.Printf("They're not the same string!")
	}
}
