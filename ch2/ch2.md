# Chapter 2. GO Basics
In this chapter, I'll cover the basics you'll need for writing programs in GO. Things like using the toolchain, data types, and some handy conversions. I will use plenty of examples in the upcoming sections. It's okay if they do not make sense; we aim for familiarity, not mastery. That said, the best way to learn any programming language is to code along to build muscle memory, so crack open your new installation, Visual Studio Code editor.

**Note:** Here is the source code in case any code causes errors or if you get a little lost

## I. The GO Toolchain

One thing that makes developing with GO so vastly different from Python or Java is that it comes with native toolchain support from the creators of the language for an easy way to manage GO codebases.

In the first chapter, we installed GO, including the toolchain. If you remember, we used 'go build -o' or 'go run main. go' to execute the hello world program, which are *commands* that are part of the GO toolchain.

We'll briefly tour the few commands you use most when writing GO.

### I.a 'go mod'

Go mod is used for dependency management. It makes it such that when you download another GO project, you'll be guaranteed that the code will work and not be broken due to dependencies.

Typically, you'll only need to run it every once in a while.

So here are the two commands you'll need most.

#### go mod init

You use `go mod init` to declare that the code you're working on is a *module* that requires x,y, and z dependencies; the dependencies will be shown within the `go.mod` that is created.

Typically, you use the format

`go mod init github.com/<yourUserName>/nameOfProject`.

#### go mod tidy

Go mod tidy eliminates dependencies that aren't used so that your codebase stays light.

To use it run,

`go mod tidy`

### I.b go get

So, GO requires that the dependencies for a project be installed locally on your machine as you develop and if they aren't then you'll see an error like:

`ERROR github.com/pkg/name is not installed.`

And to solve run,

`go get github.com/pkg/name`

But the catch is that you must have done `go mod init` first.


### I.c 'go fmt'

There are many heated arguments on the internet about how a programming language's codebase should look in terms of format.

But with GO, there is only one option, and you can use it immediately.

Imagine, for some reason, that you can't access Visual Studio Code, or the extension stops working, and the go code you wrote from [chapter 1. Setting Up Developer Environment]() looks like this:

```
package main 
import "fmt"
func main(){
fmt.Println("Hello, world!")
}
```

Then to make it GO-compliant, all you'd need to do is,

```
go fmt main.go
```
And then your code will look like this.

```
package main
	 
import "fmt"
	 
func main() {
	fmt.Println("Hello, world!")
}
```
Although having a linter like the GO extension installed is preferred to this method, it is good to know for odd occasions when you don't have it.

### I.d 'go run'

Compiles and runs your GO program all at once very useful when you're working on new code and want instant feedback without having to do 'go build' first.

```
go run main.go
```


### I.e 'go build'

You'll run this when you've finished your program and want to use it as an executable. When you're done building a Command Line Interface (CLI) tool or a Web App with a server component, you'll want executables.

And what's better is that with `go build`, you can specify which operating systems you want it to run on.

`go build -o hello env GOOS=target-OS GOARCH=target-architecture`

Don't worry if this doesn't make sense right now. Just note that when you want to take your program executable to a different operating system, it is super easy to do by using GOOS and GOARCH *environment variables*.

### I.f 'go test'

Testing your code speeds up the development process and is often a requirement in the professional world, so having support right out of the box is fantastic.

Testing is a complete topic worthy of a book or blog post. If you want to look for an in-depth explanation of it, then read my post [GO Test Driven Development](https://www.toul.io/golang-tdd/)

Otherwise, know that if you see *_test.go* files within a codebase that you can run

`go test . `

To run the tests for that codebase.

## II. Data Types

Computers, unlike humans, need precise instructions. Part of the specificity of the instructions is what data type it is to process.

So, we'll cover the significant data types in the next section that you'll most likely use and encounter while working with GO.

### II.a Integers

Integers are your whole numbers, meaning there are no decimal points.

You use the *signed* integer type if a number can be negative. Otherwise, the *unsigned* type should be used.

#### Signed

Imagine you have a program to analyze your bank account balance after paying monthly bills, and you've unfortunately had to think about a case where you are negative money.

You've overspent and might need a credit card, loan, or part of the next paycheck to pay the remainder.

So, in this case, your result could be *negative*, which tells us to use a signed integer.

First, create a new directory, 'ex1', and then create a new file named 'main.go'.

And type the following.

```
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

```

then run

`go run main.go`

And you'll see you indeed might have a problem.

Now that you know when you might use signed integers, it is worth *mentioning* that Go lets you specify how many bits an integer can contain with int8,int16, int32, and int64, but for most of the readers, **int** will more than get the job done.

#### unsigned

Unsigned integers will not be negative, so if we take the same code and swap 'int' with 'uint', you'll see that it breaks.

```
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

```

It outputs a nonsensical "18446744073709550616" instead of '-1000' as in the previous example.

Just like with 'int', there are uint8, uint16, uint32, and uint64 versions, and just like with 'int' most readers will only need 'uint.'

### Floats

Floats are your decimal numbers; by default, they can handle both positive and negative numbers. In general, I'd use a float for the calculation above because very rarely is anything, only whole numbers.

I recommend something other than the float if you're on a project where memory is a concern or if you know that decimals will not appear.

```
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

```
However, there is no 'float' like with 'uint' or 'int', so you must declare either 32bit or 64bit-sized floats. Don't worry float32 will handle most of your use cases.

> An explanation about bits is outside this book's scope, and it isn't necessary to write GO code to automate your daily tasks.

### II.b string

The data you're reading in this book and often contains within documents is of the 'string' data type. And chances are this is the most common data type you'll be using and seeing.

In GO and many other languages, string types are represented with quotation marks. Like so

```
// ex4
package main

import "fmt"

func main() {
   var example string = "Hello, world!"
   fmt.Printf("%s", example)
}

```

### II.c booleans

Boolean variables are either **true** or **false**, you'll use this when you're checking for conditions in a program, and we'll use them later.

```
// ex5
package main

import "fmt"

func main() {
   //Default value will be false
   var a bool
   fmt.Printf("a's value is %t\n", a)
}

```

## III. String Concatenation

String concatenation is the idea of adding two strings together to make one string, and as such, it uses the '+' symbol.

```
// ex6
package main

import "fmt"

func main() {
   var s1 string = "Hello "
   var s2 string = "world!"
   fmt.Printf("%s", s1+s2)
}

```

## IV. Variables

We've used variables throughout the examples, as shown by the 'var' keyword. Think of a variable like an empty box that you have left over from moving.

Before you move, you put stuff inside of it, and if you are of the well-organized type, you might even have the same type in one box.

For example, if packing a box for the kitchen, you might have a box labeled 'kitchen-glasses', and within all glasses from your cabinets.

So, only one data type goes into the variable, as only glasses go into the kitchen-glasses box.

And this is how GO expects you to use variables in a very organized manner.

However, there is a trick that you can use to avoid having to type 'var' and the type the data is explicitly, and that's the ':=', which is called the **short declaration operator**. But you can only do it within *functions*.

Here's what I mean:

```
// ex7
package main

import "fmt"

func pay(money, bills int) int {
   return money - bills
}
func main() {
   // var, int -> :=
   paycheck, bills := 4000, 5000
   fmt.Printf("bank account balance: %d\n", pay(paycheck, bills))
}

```

So, the 'main()' is a 'func'; hence, we can use the ':=', which is generally the preferred way to save on words. In codebases, less is more.

### IV.a Assignment statements

Unlike the ':=' that we previously discussed, an assignment statement uses '=', which we used in all the examples, and it requires us to use 'var' and the data type after the variable name.

E.g.

```
// ex8
package main

import "fmt"

func main() {
   var example string = "Gentle reminder"
   fmt.Printf("%s", example)
}

```
In this, we are saying that **example** is a variable of the typed string, and we are *assigning* the string "Gentle reminder" to be stored within it.

### Naming Conventions

Typically, GO programmers prefer short, concise variable names with no '_' or '-', called camelcase and kebob case, respectively.

So, you'll see terms for variables being as few characters as possible and function names like so:

```
// ex9
package main

import "fmt"

func getUserName() string {
   return "Reader"
}
func main() {
   n := getUserName()
   fmt.Printf("Hello Dear %s", n)
}

```
Naming things like this is the opposite of Python's preference for names like 'name_of_user = get_user_name(),' which encourages spelling things out as much as possible.

## V. Type Conversions

Imagine you have data in a document of '1.99' and want to use it as a float for some calculation. Sometimes, you'll find variables and want to change them to another type. Doing so is called casting or conversion in programming speak, and it is probably most familiar with **string** data.

Here are a few to keep in mind as you work with GO.

### golang string to int
```
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

```


### int to string
```
// ex11
package main

import (
   "fmt"
   "strconv"
)

func main() {
   i := 10
   s1 := strconv.FormatInt(int64(i), 10)
   s2 := strconv.Itoa(i)
   fmt.Printf("%v, %v\n", s1, s2)
}

```
### golang bytes array to string
```
// ex12
package main

import "fmt"

func main() {
   s := string([]byte{65, 66, 67, 226, 130, 172})
   fmt.Println(s)
}

```
### golang bool to string

```
// ex13
package main

import (
   "fmt"
)

func main() {
   B := true
   str := fmt.Sprintf("%v", B)
   fmt.Println(str)
}

```

## VI. Comparisons

They are rarely used alone and typically go with control-flow or conditional statements, which we'll go over later, but I'm showing them here to build familiarity.

### equal to
Compares whether or not two values stored in variables are equal.

```
// ex14
package main

import (
   "fmt"
)

func main() {
   fmt.Println("let's see if these two identical strings evaluate to equal")
   s1 := "hi"
   s2 := "hi"
   // okay, I showed a little conditional flow here
   if s1 == s2 {
      fmt.Printf("yes, they are!")
   }
}

```
### not equal to

```
// ex15
package main

import (
   "fmt"
)

func main() {
   fmt.Println("let's see if these two non-identical strings evaluate to equal")
   s1 := "hi"
   s2 := "these are not the droids you're looking for"
   // okay, I showed a little conditional flow here
   if s1 != s2 {
      fmt.Printf("They're not the same string!")
   }
}

```
### greater than

```
// ex16
package main

import (
   "fmt"
)

func main() {
   bills := 13642.32
   paycheck := 10000.00
   if bills > paycheck {
      fmt.Printf("Maybe I should update my resume...")
   }
}


```

### less than
```
// ex17
package main

import (
   "fmt"
)

func main() {
   paycheck := 13642.32
   bills := 10000.00
   if bills < paycheck {
      fmt.Printf("alright, alright, alright")
   }
}

```

### greater than or equal to

```
// ex18
package main

import (
   "fmt"
)

func main() {

   bills := 13642.32
   paycheck := 10000.00
   if bills >= paycheck {
      fmt.Printf("I'll never gonna financially recover from this")
   }
}

```

### less than or equal to

```
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

```





