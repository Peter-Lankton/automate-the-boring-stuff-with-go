// ex.1 repetitive code
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Add two numbers..")

	var a int
	var b int

	for {
		fmt.Println("Enter your number: ")
		fmt.Scan(&a)
		fmt.Println("Enter your number: ")
		fmt.Scan(&b)
		fmt.Println("result ", a+b)

	}
}
