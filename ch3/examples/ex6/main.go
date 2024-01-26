// ex 6. If you're a reader, then a thousand thanks
package main

import "fmt"

func main() {
	fmt.Println("Enter your name")
	var n string
	fmt.Scan(&n)
	switch {
	case n == "Toul":
		fmt.Printf("The author is %s ", n)
	default:
		// a thousand thanks if you're not the author
		for i := 0; i < 1000; i++ {
			fmt.Printf("\ni: %d Thank you %s for reading", i, n)
		}
	}
}
