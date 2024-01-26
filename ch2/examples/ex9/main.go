// ex9
package main

import "fmt"

func getUserName() string {
	return "Reader"
}
func main() {
	n := getUserName()
	fmt.Printf("Hello Dear %s", n)
}
