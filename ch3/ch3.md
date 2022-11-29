# Ch. 3 Control Flow



In the previous chapter, we covered the basics of GO data types and, in a few examples, used flow control to print messages.

In this chapter, we'll use what we learned about comparison operators and boolean values to learn more about flow control.

## Elements of Flow control

Programs are like water in a river because they tend to flow in one direction. And like water in a river, parts of it can go off into separate outlets based upon the river's conditions, such as a boulder in the stream or a dam.

However, unlike water in a river, it's much easier to divert the flow of a program.

To do so, we use conditions to evaluate a variable's value or the output of a function and execute a different code block.

Let's use an example with comments to see.

We are using the familiar 'if' from the previous chapter.


```
// Ex. 1 Tell me if a number is even or odd
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// use the rand library to generate a random integer.
	n := rand.Intn(100)
	// 'if' condition
	if n%2 == 0 {
		// code block that's executed if the condition is true
		fmt.Printf("Number is even: %d", n)
	}
	// code block that's executed if it isn't true
	fmt.Printf("Number is odd: %d", n)
}
```

In this program, we're checking whether a *'random'* integer number is even or not via the *modulo operator* `%`, which is a fancy way of saying 'dividing with remainders'.

If you recall, an *even* number can be divided in two without any remainders, whereas an *odd* number cannot.

Don't worry about the math. Just note that the logic of the program changes based on the condition, which is the *'random'* value of the **n**  variable.

Now, let's introduce the rest of the gang for flow control.

## Flow Control Statements

## 'if' statements

These are the most used control statements, and you'll almost always use them when adding 'logic' to your programs.

The general pattern is that your program has a chunk of code that will only be executed if the condition is considered 'true' or 'false.'



## 'else' statements

Technically, in the previous 'if' example, we could've had an 'else' statement to print the 'number is odd:' part, but we kept it cleaner by leaving it out, which is a preference amongst programmers.

```
// ex. 2 else
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// use the rand library to generate a random integer.
	n := rand.Intn(100)
	// 'if' condition
	if n%2 == 0 {
		// code block that's executed if the condition is true
		fmt.Printf("Number is even: %d", n)
	} else {
		// code block that's executed if it isn't true
		fmt.Printf("Number is odd: %d", n)
	}

}
```

Else statements say 'hey, if this piece of logic turns out to be false or something else, then do this.'

Typically, 'else' statements aren't preferred because they aren't explicit and will catch only when the condition is false.

Let's introduce the else's, cousin.

## else if

Here, we can communicate what we want to happen 'if' the first condition evaluates to 'false' and explicitly define what should happen.

```
// ex. 3 else if 
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// use the rand library to generate a random integer.
	n := rand.Intn(100)
	// 'if' condition
	if n%2 == 0 {
		// code block that's executed if the condition is true
		fmt.Printf("Number is even: %d", n)
	} else if n%2 == 1 {
		// code block that's executed if it isn't true
		fmt.Printf("Number is odd: %d", n)
	}

}
```

## switch

The 'switch' control flow works much the same as using 'if' 'else' or 'else if' but presents differently in the form of 'cases'. A 'case' is the same as a condition and may be preferred by those with a mathematical background. Either style works; use whichever makes sense to you.

```
// ex. 4 switch
package main

import (
	"fmt"
	"time"
)

func main() {

	// You can use commas to separate multiple expressions
	// in the same `case` statement. We use the optional
	// `default` case in this example as well.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// `switch` without an expression is an alternate way
	// to express if/else logic. Here we also show how the
	// `case` expressions can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}

```

In the above, we say 'switch' based upon the following 'case' and pass in some data. The same again can be accomplished by using 'if' and 'else if'

```
// ex. 5 convert switch to if else
package main

import (
	"fmt"
	"time"
)

func main() {
	if t.Hour() < 12 {
		fmt.Println("before noon")
	} else if t.Hour() > 12 {
		fmt.Println("it's after noon")
	}
}
```
In general, I prefer using 'switch' and including the optional 'default' case because I find it easier to test my 'cases' and to convey to others what I'm trying to do with my code.

## Loops

Now that we have some logic to control a program's flow, it's time to introduce 'loops'. A loop is a way we tell computers to "do work for x amount of times". Let's combine it with loops to make a more intelligent program.

A 'for' loop also uses a condition to do its work, but it works 'until' that condition is satisfied. So, for example, the saying 'a thousand thanks' we'll have a program that will say 'thank you' a thousand times based upon a reader's name.

```
// ex 5. If you're a reader, then a thousand thanks
package main

import "fmt"

func main() {
   fmt.Println("Enter your name")
   var n string
   fmt.Scan(&n)
   switch {
   case n == "Toul":
         fmt.Printf("The author is %s ",  n)
   default:
     // a thousand thanks if you're not the author
      for i := 0; i < 1000; i++ {
         fmt.Printf("\ni: %d Thank you %s for reading", i, n)
      }
   }
}

```
In the above the 'for' loop syntax is *for* **initial condition**; *compare* ;**if not met then do another** { // work to be done }

> In most programming languages, it is common to start at '0', so keep this in mind when working with arrays, slices, and initial loop conditions.

Okay, let's have some fun by building a few small games to show your friends and family at the winter holiday parties.

Don't worry if some syntax is unfamiliar. We want to focus on building muscle memory and familiarity with GO. So, try your best to copy the code and get in running. If you have any troubles at all the [source code is available](https://github.com/toul-codes/automate-the-boring-stuff-with-go/tree/main/ch3)

## Guess the Number

In this first game we'll have the computer pick a random number and then try to guess it by inputting guesses, but the catch is we only get '3' attempts!

But, we'll give the user clues to increase the chances of winning within 3.

```
package main

import (
   "fmt"
   "math/rand"
   "time"
)

func main() {
   fmt.Println("Game: Guess a number between 0 and 10")
   fmt.Println("You have three(3) tries ")

   source := rand.NewSource(time.Now().UnixNano())
   randomizer := rand.New(source)
   secretNumber := randomizer.Intn(10)

   var guess int

   for try := 1; try <= 3; try++ {

      fmt.Printf("Attempt: %d\n", try)
      fmt.Println("Enter your number: ")
      fmt.Scan(&guess)
  
      switch {
      case guess < secretNumber:
         fmt.Printf("Sorry, wrong guess ; number is too small\n ")
      case guess > secretNumber:
         fmt.Printf("Sorry, wrong guess ; number is too large\n ")
      case guess == secretNumber:
         fmt.Printf("You win!")
      case try == 3:
         fmt.Printf("Game over!!\n ")
         fmt.Printf("The correct number is %d\n", secretNumber)
      default:
         fmt.Println("I don't recognize that input. Try a number like '1'")
      }
   }
}
```

## Advanced: Rock, Paper, Scissors

Now, for a more fun and more advanced example to wet your programmer appetite.

```
package main

import (
   "fmt"
   "math/rand"
   "strings"
   "time"
)

func getInput() string {
   // get input and use strings package to convert to lowercase b/c this is a good practice when working with inputs 
   // in general
   fmt.Print("Pick [r]ock, [p]aper, or [s]cissors:   ")
   var input string
   fmt.Scanln(&input)
   return strings.ToLower(input)
}

type Move int

const (
   Rock Move = iota
   Paper
   Scissors
)

var validInput = map[string]Move{
   "r":        Rock,
   "rock":     Rock,
   "p":        Paper,
   "paper":    Paper,
   "s":        Scissors,
   "scissors": Scissors,
}

var inputs = [...]string{
   Rock:     "Rock",
   Paper:    "Paper",
   Scissors: "Scissors",
}

func checkInput() Move {
   // start a loop until we get valid input
   for {
      // Does the user input match the map validInput?
      if move, ok := validInput[getInput()]; ok {
         fmt.Println("Player Chooses", inputs[move])
         return move
      }

      fmt.Println("I didn't understand your choice, please retry")
   }

}

var compare = [...]string{
   Rock:     "You Tied.",
   Paper:    "You Lost.",
   Scissors: "You Win!",
}

func main() {
   // for {} says do this until the user wants to stop with ctrl+c which is the universal way to kill a program 
   // that's running in the terminal.
   // this is also known as 'while' loop in other languages and is how video games work (while user playing keep 
   // running game engine) 
   for {
      rand.Seed(time.Now().Unix())

      // Get Valid Player Input
      userChoice := checkInput()

      // Randomly Assign Computer Choice
      computerChoice := Move(rand.Intn(3))

      computerInput := inputs[computerChoice]

      fmt.Printf("Computer chooses: %s\n", computerInput)

      // Check to see who won
      fmt.Printf("Result: %s\n",
         compare[((computerChoice-userChoice)%3+3)%3])
   }

}


```