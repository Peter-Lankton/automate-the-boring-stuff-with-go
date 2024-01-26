

You've already seen **func**'s multiple times so far, but let's take this chapter to understand them better and when to use them when automating work.

Functions are reusable work units for a program to use, and they save programmers from repeating themselves.

> You might hear this paradigm as **D.R.Y.** when speaking with other coders, which stands for "Don't repeat yourself".

Functions are like small programs within your program. And GO, in particular, must always have the **func main() {}** within the `main.go` file.

The **main** func is to tell the compiler (go build) that this is where the code starts. Remember, computers execute code line by line, so it helps to tell them which line to begin with.

## I.Parts of a function

Let's start with a basic function and try to understand what each part of it is called and what it is responsible for.

```
// ex. 1
package main 

func NameIs(name string){ // (1)
	fmt.Printf("Your name is: %s", name)
}

func main() {
	NameIs("Gopher") // (2)
}
```

**(1)** The characters succeeding the 'func' keyword give the name to the 'func'.

In this example, it is correct to say the function's name is, well, **NameIs**.

And the characters within the parentheses 'name string' is the accepted argument(s).

> **Note:** There can be more than one argument that a function can to do work on.


**(2)** It Is known as the 'function call,' or the function is being 'called/invoked'. That is a fancy way of saying, "the function is being used".

## I.a When to get func-e (funky)?

Frequently whenever you're writing code, you might get it working-- and that is fine and a noble goal on the first pass. However, on the second pass (immediately, once it is working), it's time to **refactor**, which means doing things like reducing repetition by using functions.

### I.b Example

Suppose we say hello to a bunch of readers without using a function. Here's how repetitive it will look.


```
// ex. 2
package main 

func main() {
	fmt.Printf("Hello %s", "Gopher")
        fmt.Printf("Hello %s", "Gopher2")
	fmt.Printf("Hello %s", "Gopher3")
        fmt.Printf("Hello %s", "Gopher4")
	fmt.Printf("Hello %s", "Gopher5")
        fmt.Printf("Hello %s", "Gopher6")
	fmt.Printf("Hello %s", "Gopher7")
        fmt.Printf("Hello %s", "Gopher8")
	fmt.Printf("Hello %s", "Gopher9")
        fmt.Printf("Hello %s", "Gopher10")

}
```
Let's do a **refactor** to make it a bit **DRY**er.

```
// ex. 3
package main

import "fmt"

func greetReaders(n []string) {
	for _, name := range n {
		fmt.Printf("\nHello %s", name)
	}
}

func main() {
	rdrs := []string{
		"Gopher",
		"Gohper2",
		"Gopher3",
		"Gohper4",
		"Gopher5",
		"Gohper6",
		"Gopher7",
		"Gohper8",
		"Gopher9",
		"Gohper10",
	}
	greetReaders(rdrs)

}
```

## II. Multiple different kinds of arguments

Suppose we wanted to print a reader's name and age. Then we'll need string type data and number type data. In GO the arguments can be a mix of different types like string and int, not just a mono-type.

```
// ex. 5 Multiple Argument types
package main

import "fmt"

func HappyBirthday(n string, age int) {
	fmt.Printf("Happy birthday %s, You're turning %d", n, age)
}

func main() {
	HappyBirthday("Go", 20)

}
```

## III. `return` values

So far, the functions defined haven't used the 'return' keyword, which means they're only doing something with the data passed in but not giving it back.

Let's go back to our even/odd examples and sort integer arrays into odd and even.

```
// ex. 4 Returning values
import "fmt"

func EvenOdd(n []int) ([]int, []int) {  // (1)
	var odd, even []int
	for _, num := range n {
		if num%2 == 0 {
			even = append(even, num)
		}
		if num%2 == 1 {
			odd = append(odd, num)
		}

	}
	return odd, even // (2)
}

func main() {
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	odd, even := EvenOdd(num) // (3) 
	fmt.Println("Odd: ", odd)
	fmt.Println("Even: ", even)

}
```
**(2)** The 'return' keywords say at the end of this function call, give the *value* to the caller. The big idea here is that the 'value' returned can be 'assigned' to variables, which can then be passed further along in the program.

Notice, in GO, that you can return multiple values and assign them at once in **(3)**. However, this can be tricky, so it's not recommended as you're starting out.

**(1)** The astute reader probably caught the additional parentheses following the accepted arguments; ([]int, []int). The second set of parentheses is required when returning multiple values. It lets the reader know the `types` expected after a successful invocation of the func.

Similarly, you can return multiple different data types as well.

```
// ex. 5 Returning multiple values 
package main

import "fmt"

func diffVals(n string, age int) (string, int) { // (1)
	return n, age
}

func main() {
	name, age := diffVals("Gopher", 20)
	fmt.Printf("name: %s, age: %d", name, age)
}
```

**(1)** Shows that a string type and int type will be returned at the end of a successful function call.

As you write code to solve problems, your programs will be more extensive than 20 lines.

Having functions that return data makes it simpler to keep track of what's going on in your program for you and other readers.

## IV. Local variables and Global Variables

In all our code so far, we've mainly been using `local` variables, and now it is time to discuss what that means and what `global` means.

The distinction is that a `local` variable and its value are only accessible within the block of code in which it exists. Whereas a `global` variable may be accessed outside of the code block and across all of your program files. In general, `global` variables are frowned upon in GO programs, but in other languages, they're not so frowned upon.

Let's see an example.

```
// ex. 6 Globals are tricky
package main

import "fmt"

// global variables
var Author = "Gopher"
var Age = 20

func LocalAuthor(name string, age int) {
	Author = name
	Age = age
	fmt.Printf("\nAuthor: %s, \nAge: %d", name, age)
}

func main() {
	LocalAuthor("Gopher2", 40)
	fmt.Printf("\nAuthor: %s, \nAge: %d", Author, Age)

}
```

Above, the `global` variables have their values reassigned within the function call of `LocalAuthor`, which might cause problems in other parts of the program if the expected values were "Gopher" and `20`.

To get the expected behavior, define local variables within the function 'n' and 'a' for name and age, respectively.

```
// ex. 7 Locals are preferred
package main

import "fmt"

// global variables
var Author = "Gopher"
var Age = 20

func LocalAuthor(name string, age int) {
	n := name // don't reference the global variables
	a := age
	fmt.Printf("\nAuthor: %s, \nAge: %d", n, a)
}

func main() {
	LocalAuthor("Gopher2", 40)
	fmt.Printf("\nAuthor: %s, \nAge: %d", Author, Age)
}

```

## V. Error handling

One of the few areas where GO receives complaints is with error handling in that it is often verbose and redundant. However, the idea is that errors should be explicitly handled and error situations described.

Let's set up a function that will return an 'error' type if the received argument is 'empty'.

```

package main

import (
	"errors"
	"fmt"
	"log"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
      // If a name isn't received then there's no one to greet and it should return an error
	if name == "" {
		return "", errors.New("Err: empty name")
	}
        // otherwise great the person and return a nil for error
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := Hello("")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
```

Error handling is a must if you're going to be building functions that you intend to share with the world or a team or with your future self.

However, just starting out it might not be as necessary when you're hacking something together.

But, now you are aware of the general style if you should be happen to digging around through some package documentation.

## Conclusion

Functions are an essential paradigm in any program language-- in fact there's languages that are purely **functional**. The key takeaways are that as a reader of a function you'll probably be most interested in what arguments it takes and what values it returns, not so much the code is within the function.



