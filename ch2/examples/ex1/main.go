// ex1
package main

import "fmt"

func pay(money, bills int) int {
	return money - bills
}

func main() {
	var paycheck, bills int = 4000, 5000
	fmt.Printf("bank account balance: %d\n", pay(paycheck, bills))
}
