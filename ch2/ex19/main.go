// ex19
package main

import (
	"fmt"
)

func main() {
	paycheck := 13642.32
	bills := 10000.00
	if bills <= paycheck {
		fmt.Printf("seems risky.")
	}
}
