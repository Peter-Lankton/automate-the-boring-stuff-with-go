Slices in Golang are the same as dynamic arrays. An array is the first data structure you'll probably learn about in any book related to learning a programming language. Slices can contain any data type, including a slice itself, making arranging data into hierarchical structures simpler.
In this chapter, we'll go over the basics of Slices, and I'll teach you about some of the most useful methods available to slices.

## The Slice Data Type

The slice type has an upper and lower bound separated by a colon but is not inclusive, meaning that the upper bound value isn't used.
For example

```
package main

import "fmt"

func main() {
	evens := [5]int{2, 4, 6, 8, 10 }
    var s []int = evens[1:4]
    fmt.Println(s)

}

```

will output
"4 6 8 "

## Accessing values in a slice with Indexes

Like many other languages in GO, getting a single value from a slice is possible. Let's say instead of using evens[1:4], it had read evens[0]  will only output "2".
Slice of Slices
Slices may contain any type, including other slices, which can be considered a grid when used. A familiar example for those not into math is Tic-tac-toe rather than the x,y coordinate plane.

```
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

```

Which will output,

```
X _ X
O _ X
_ _ O
```

## Using empty slices
An empty-sized slice is helpful when you don't know how many elements will be in it, but you do plan on using it. So you'd declare it like this.

```
package main

import "fmt"

func main() {
	a := []int{}
	fmt.Println(a)

}

```

Which will output '[],' showing that it's an empty slice waiting to be filled.

## Appending to a Slice
When working with slices, it is often necessary to add new elements to the slice, and the appropriate way of doing so in go is to use the 'append' method.

```
package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```


## Getting a Slice's length with len() func.
Similarly, to add elements to a slice, it is often essential to know how many features are inside a slice.

```
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(len(s))


}
```

## Changing Values in a Slice with Indexes
Sometimes, you'll also want to update a specific value in the Slice without altering others.

```
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	s[1] = 234234
    fmt.Println(s)


}
```

Which will output`[2 234234 5 7 11 13]`

## Using Slices
Okay, now that the basics of slices are down, let's see how and why we'd probably use a slice to automate our mundane tasks.

### For loops
Loops and Slices go together like milk and cookies, and GO a unique range syntax for going through a slice, element by element and value by value.

Suppose you had data about the prices of paper you were selling and needed to update them to account for inflation.

```
package main

import "fmt"



func main() {
	paperPrices := []float32{1.00, 2.00, 4.00, 8.00, 16.00, 32.00, 64.00, 128.00}
	inflatedPaperPrices := make([]float32, len(paperPrices))
	for i, v := range paperPrices {
		inflatedPaperPrices[i] = (v * 0.08) + v
	}
    fmt.Println(inflatedPaperPrices)
}
```


Notice that the 'make' keyword was used when declaring the inflatedPaperPrices slice. The make statement is used in GO to create variables with predetermined lengths and types. It is often best practice and, in this case, must be done.

Let's break it down, starting with 'i' and 'v'– these are the for loop variables that only exist within the for loop and may be named anything, but typically 'i' and 'v' are used by tradition. Because 'i' can be thought of as 'index' and 'v' as 'value'.

So, range lets us iterate (go through) a slice and access either the value or index or both.  The range syntax is a shortcut to write a for loop in GO. The code above could be written in a longer form like so.

```
package main

import "fmt"



func main() {
	paperPrices := []float32{1.00, 2.00, 4.00, 8.00, 16.00, 32.00, 64.00, 128.00}
	inflatedPaperPrices := make([]float32, len(paperPrices))
	for i:=0; i < len(paperPrices); i++ {
		inflatedPaperPrices[i] = (paperPrices[i] * 0.08) + paperPrices[i]
	}
    fmt.Println(inflatedPaperPrices)
}
```


## Conclusion
Now that the basics of slices have been covered, let's move on.