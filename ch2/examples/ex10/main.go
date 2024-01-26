// ex10
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	strVar := "100"
	intVar, err := strconv.Atoi(strVar)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
}
