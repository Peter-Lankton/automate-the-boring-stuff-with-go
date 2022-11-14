// ex3
package main

import "fmt"

// int -> uint -> float32 if you're coding along
func pay(money, bills float32) float32 {
	return money - bills
}
func main() {
	// int -> uint -> float32
	var paycheck, bills float32 = 4000, 5000.0
	fmt.Printf("bank account balance: %f\n", pay(paycheck, bills))
}
