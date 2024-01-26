package main

import (
	"fmt"
	"os"
	"regexp"

	"flag"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s -h\n", os.Args[0])
	} else {
		pass := flag.String("p", "", "get password")

		flag.Parse()

		regStr, _ := regexp.Compile(`([0-9a-zA-Z]){11,}`)

		if regStr.MatchString(*pass) {
			fmt.Println("Password adheres to company policy")
		} else {
			fmt.Println("Password does not adhere to company policy")
		}

		os.Exit(0)
	}
}
