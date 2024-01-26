# Chapter 7 

Think of all the times you've used **cmd + f** or **ctrl + f** on your computer to search for a phrase or string on a 
webpage or huge document.

It's also possible to do the same in programming but rather than matching a word character for character we can search for
the pattern of the string. 

Here's an example, suppose you've been handed a 1 GB document and your task is to find all the e-mails within it to build
a database for your job for whatever reason.

So you try what you might know, and that's searching for 'name@company.com' while that will yield some results it won't 
find results that might look like 'name@company.two.com', and so on.

But you know in general that an e-mail contains the char '@' and a string in front of it and a string behind it, typically
ending in '.com'. You could try to think of all the combinations but that would probably take way too long and the work
is due by end of day.

Enter *regular expressions* or commonly known as *regex* in programming.  It is a useful piece of knowledge that you'll
need more often than not.

In this chapter you'll see how to find patterns without regex and with regex, sometimes you really do need to find only a 
particular string or sub-string. Plus it is easier to start with before reaching for regex as the pattern building can 
be tricky if it isn't something you use every day.

## Finding Strings and Substrings without Regex

First, I'll clarify on what a substring means since it has been used several times so far in this chapter. A sub string 
is a smaller part of a string. 

Example:
```go
package main 
import "fmt"
func main() {
    s := "continental"
	fmt.Printf("the %s is a John Wick universe related spin off",s)
}
```
then a substring of it is 

```go
package main 
import "fmt"
func main(){
	ss := "continent"
	fmt.Printf("%s is a substring of 'contiental'", ss)
}
```

Now, that is out of the way. Let's see how to find substrings within strings. 

### Search

The GO programming language standard library has *search* built right in and this is a decent starting point for finding
strings.

```go
package main 
import (
	"fmt"
	"strings"
)
func main() {
	doc := `
        In the fast-paced world of corporate business, the pursuit of optimal efficiency is paramount. 
        Companies must constantly strive to fine-tune their operational processes, fostering an environment of maximum productivity. 
        This text aims to delve into the intricacies of optimizing synergistic business operations, highlighting the various 
        mundane aspects of this essential endeavor.To begin, it is essential to emphasize the importance of aligning organizational 
        goals with strategic planning. Companies need to set clear objectives and devise comprehensive strategies to achieve them. 
        By ensuring that these objectives are communicated effectively and consistently across all departments, companies can 
        expect to see improved coordination and coherence in their operations. Once the strategic framework  a fundamental but 
        tedious undertaking that requires careful planning, continuous process improvement, meticulous monitoring, effective time 
        management, and resource allocation. While these aspects of corporate business may not be the most captivating, they are 
        integral to the success of any organization. A relentless commitment to these seemingly banal activities is what sets the 
        stage for a successful, efficient, and sustainable corporate operation.`
	
	if strings.Contains(doc, "aspects") {
		fmt.Println("Yes, 'aspects' is in that corporate word salad")
    }
	
	if !strings.Contains(doc, "orange") {
		fmt.Println("silly Gopher there is no 'orange' in that corporate word salad")
    }
	fmt.Printf("how many times is the word 'companies' used? ans: %v", strings.Count(doc, "companies"))
	
}


```

### Replace

Sometimes when searching for a word, it is because you want to replace it. So let's take a moment to view how to do so 
in GO before getting into Regex.

```go
package main 
import (
	"fmt"
	"strings"
)
func main() {
	doc := `
        In the fast-paced world of corporate business, the pursuit of optimal efficiency is paramount. 
        Companies must constantly strive to fine-tune their operational processes, fostering an environment of maximum productivity. 
        This text aims to delve into the intricacies of optimizing synergistic business operations, highlighting the various 
        mundane aspects of this essential endeavor.To begin, it is essential to emphasize the importance of aligning organizational 
        goals with strategic planning. Companies need to set clear objectives and devise comprehensive strategies to achieve them. 
        By ensuring that these objectives are communicated effectively and consistently across all departments, companies can 
        expect to see improved coordination and coherence in their operations. Once the strategic framework  a fundamental but 
        tedious undertaking that requires careful planning, continuous process improvement, meticulous monitoring, effective time 
        management, and resource allocation. While these aspects of corporate business may not be the most captivating, they are 
        integral to the success of any organization. A relentless commitment to these seemingly banal activities is what sets the 
        stage for a successful, efficient, and sustainable corporate operation.`
	
	newDoc := strings.Replace(doc, "Companies", "Organization", 3)
	if strings.Contains(doc, "Organization") {
		fmt.Println("Organization is not present, searching...newDoc")
		if strings.Contains(newDoc, "Organization"){
			fmt.Println("Replacement successful.")
        }
    }
}
```

Now that the basics are covered let's move onto regex.

## Regex

GO comes with the 'regexp' package from the devs, so there's no need to add a third party dependency. In general, for most 
devs regexes are challenging to build.

So, here's a tool that is helpful for when trying to build a regex pattern in code [RegExr.com](https://regexr.com/) in 
it, you can practice building your regex pattern with real time feedback, as well as reference common patterns.

Now, Let's dive into some working practice examples to better illustrate.

### First Match

Using regex in GO requires `re := regexp.MustCompile(<regex-pattern-here>)` and then searching for the pattern takes a second 
line, `re.FindString(<string-to-search>)`

Let's look at some code to see it in action.

Suppose, we want to find the first 'foo' in a string discussing the various football's around the world. 

```go
package main 
import (
	"fmt"
	"regexp"
)
func main() {
    re := regexp.MustCompile(`foo.?`)
    fmt.Printf("%q\n", re.FindString("People like to tailgate for the American sport football not to be confused with football.")) // "food"
    fmt.Printf("%q\n", re.FindString("this string doesn't have the substring. Bummer..."))         // ""
}
```

Notice, the pattern required "(`foo.?`)" it is like asking where does 'foo' first appear. 

### Location

Now, suppose we're trying to find a location where treasure is buried then the syntax is mostly similar but notice that 
we dropped the '.' from the MustCompile expression. 

```go
package main 
import (
	"fmt"
	"regexp"
)
func main() {
    re := regexp.MustCompile(`treasure?`)
    fmt.Println(re.FindStringIndex("Some maps lead to riches and some maps don't. But this map if you rotate it and squint your eyes, X marks the treasure."))    // [1 3]
    fmt.Println(re.FindStringIndex("there be none here ya know try another map") == nil) // true
}
```

It is the subtleties, of regex that make it challenging to work with as one char change can be often what is needed, so
it is best to test and change to figure out what the exact pattern you need is.


### All Matches 

Instead, of finding a treasure, maybe you'll need to find all the matches of a substring in a document instead.

```go
package main
import (
	"fmt"
	"regexp"
)
func main() {
    re := regexp.MustCompile(`a.`)
    fmt.Printf("%q\n", re.FindAllString("An apple in the big apple is known to cause lots of apathy for some reason", -1))
    fmt.Printf("%q\n", re.FindAllString("paranormal", 2))
    fmt.Printf("%q\n", re.FindAllString("graal", -1))  
    fmt.Printf("%q\n", re.FindAllString("none", -1))      
}
```

## Projects

Now, that we've seen a few examples let's build some projects that will show real world use cases of regex rather than 
practice examples.

### 1. Regex E-Mail finder from clipboard
In this project, we'll use Regex to determine if a piece of text copied to our clipboard contains an e-mail. The pattern 
of finding an e-mail in a string of text is handy as you'll probably receive data where it will be useful to the group 
based on a common characteristic.

If you're interested in building a web application with sign-up / sign-in input fields based on e-mails, this can be 
useful for determining whether the input is of the correct form.


**Note:** There is a third-party package dependency in this project, but I trust you remember how to handle that requirement 
*hint* go mod init & go get

```
package main

import (
	"fmt"
	"regexp"
	"strings"
	"github.com/atotto/clipboard"
)

func main() {
	// create email regexp
	regMail, _ := regexp.Compile(`[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,6}`)

	// read os buffer
	// find email regexp
	text, _ := clipboard.ReadAll()
	var mailAddr []string
	// found e-mail
	if regMail.MatchString(text) {
		mailAddr = regMail.FindAllString(text, -1)
	}

	// Print found e-mails on the terminal
	if len(mailAddr) > 0 {
		clipboard.WriteAll(strings.Join(mailAddr, "\n"))
		fmt.Println("Copied to clipboard:")
		fmt.Println(strings.Join(mailAddr, "\n"))
	} else {
		fmt.Println("No email addresses found.")
	}
}

```


### 2. Regex Password Complexity Checker
Password complexity is an additional feature you will want to have for an application. This is because simple passwords 
are easier for cyber criminals to obtain/guess.

Hence, it is wise to enforce a password complexity of at least 11 characters with some capital letters and numbers thrown in.  
So, let's build a small program to check whether a given password passes the complexity requirement.

```
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
			fmt.Println("Password ok")
		} else {
			fmt.Println("Bad password")
		}

		os.Exit(0)
	}
}

```

Running the program and passing in your possible password should result in either a Password OK or Bad Password response 
from the program:

```
t@m1 regexppass % go run main.go --p=23498aosethuaosthAT
Pass ok
t@m1 regexppass % go run main.go --p=2Aoeue             
Bad password
t@m1 regexppass % go run main.go --p=2AoeueEEEE
Bad password
```


### 3. Quiz Builder

Now, this project is the first in our queue at automating a mundane work task--your teacher friends will love you!
In this project, you'll build a quiz generator around the U.S. and its capitals. However, you will generally have the 
format of a quiz-building piece of software that can be changed to generate different sorts of quizzes. Maybe, you'll 
change it to do the capitals of your home country and its states.

```
package main

import (
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	capitals := map[string]string{
		"Alabama":        "Montgomery",
		"Alaska":         "Juneau",
		"Arizona":        "Phoenix",
		"Arkansas":       "Little Rock",
		"California":     "Sacramento",
		"Colorado":       "Denver",
		"Connecticut":    "Hartford",
		"Delaware":       "Dover",
		"Florida":        "Tallahassee",
		"Georgia":        "Atlanta",
		"Hawaii":         "Honolulu",
		"Idaho":          "Boise",
		"Illinois":       "Springfield",
		"Indiana":        "Indianapolis",
		"Iowa":           "Des Moines",
		"Kansas":         "Topeka",
		"Kentucky":       "Frankfort",
		"Louisiana":      "Baton Rouge",
		"Maine":          "Augusta",
		"Maryland":       "Annapolis",
		"Massachusetts":  "Boston",
		"Michigan":       "Lansing",
		"Minnesota":      "Saint Paul",
		"Mississippi":    "Jackson",
		"Missouri":       "Jefferson City",
		"Montana":        "Helena",
		"Nebraska":       "Lincoln",
		"Nevada":         "Carson City",
		"New Hampshire":  "Concord",
		"New Jersey":     "Trenton",
		"New Mexico":     "Santa Fe",
		"New York":       "Albany",
		"North Carolina": "Raleigh",
		"North Dakota":   "Bismarck",
		"Ohio":           "Columbus",
		"Oklahoma":       "Oklahoma City",
		"Oregon":         "Salem",
		"Pennsylvania":   "Harrisburg",
		"Rhode Island":   "Providence",
		"South Carolina": "Columbia",
		"South Dakota":   "Pierre",
		"Tennessee":      "Nashville",
		"Texas":          "Austin",
		"Utah":           "Salt Lake City",
		"Vermont":        "Montpelier",
		"Virginia":       "Richmond",
		"Washington":     "Olympia",
		"West Virginia":  "Charleston",
		"Wisconsin":      "Madison",
		"Wyoming":        "Cheyenne",
	}

	var states, capitalsItems []string
	for y, x := range capitals {
		capitalsItems = append(capitalsItems, x)
		states = append(states, y)
	}

	for i := 0; i < 35; i++ {
		// сreate the quiz text file
		str1 := "quiz_" + strconv.Itoa(i+1) + ".txt"
		quizFile, err := os.Create(str1)
		check(err)
		defer quizFile.Close()
		// create the answer key to the quiz
		str2 := "answer_key_" + strconv.Itoa(i+1) + ".txt"
		answerKeyFile, err := os.Create(str2)
		check(err)
		defer answerKeyFile.Close()

		// Create portion for students to fill out
		quizFile.WriteString("Student Number:\n\nName:\n\nDate:\n\n")
		str3 := "Quiz " + strconv.Itoa(i+1)
		quizFile.WriteString(strings.Repeat(" ", 20) + str3)
		quizFile.WriteString("\n\n")

		rand.Seed(time.Now().UnixNano())
		// mix of the States
		shuffle(states)

		// Iterate through and build the question out 
		for j := 0; j < 50; j++ {
			correctAnswer := capitals[states[j]]
			wrongAnswers := make([]string, len(capitalsItems))
			copy(wrongAnswers, capitalsItems)

			// shuffle wrong answers
			answNoCorrect := make([]string, len(wrongAnswers)-1)
			for l := 0; l < len(wrongAnswers); l++ {
				if wrongAnswers[l] == correctAnswer {
					copy(answNoCorrect, removeAtIndex(wrongAnswers, l))
				}
			}

			// create answer options A-D
			var answerOptions []string
			for l := 0; l < 3; l++ {
				answerOptions = append(answerOptions, answNoCorrect[l])
			}
			answerOptions = append(answerOptions, correctAnswer)
			shuffle(answerOptions)

			// create question
			str3 := strconv.Itoa(j+1) + " What is the Capital of " + states[j] + "?" + "\n"
			quizFile.WriteString(str3)
			strAbcd := "ABCD"
			for l := 0; l < 4; l++ {
				strAnsw := string(strAbcd[l]) + ". " + answerOptions[l] + "\n"
				quizFile.WriteString(strAnsw)
			}
			// make quiz and save it
			quizFile.WriteString("\n")

			// make answer key and save it
			strAnswerOk := ""
			for l := 0; l < len(answerOptions); l++ {
				if answerOptions[l] == correctAnswer {
					strAnswerOk += string(strAbcd[l])
				}
			}
			strCorAnsw := strconv.Itoa(j+1) + ". " + strAnswerOk + "\n"
			answerKeyFile.WriteString(strCorAnsw)
		}
	}
}

// helper functions for making quiz building easier
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func shuffle(a []string) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func removeAtIndex(source []string, index int) []string {
	lastIndex := len(source) - 1
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}

```

## Conclusion

Now, you know how to find things in an even more efficient manner for when `cmd + f` or `ctrl + f` cannot get the job 
done. 