// ex2
package main

import "fmt"

// changed int to uint
func pay(money, bills uint) uint {
	return money - bills
}
func main() {
	// changed int to uint
	var paycheck, bills uint = 4000, 5000
	fmt.Printf("bank account balance: %d\n", pay(paycheck, bills))
}
