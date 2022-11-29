// ex. 5 convert switch to if else
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	if t.Hour() < 12 {
		fmt.Println("before noon")
	} else if t.Hour() > 12 {
		fmt.Println("it's after noon")
	}
}
