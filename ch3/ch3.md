# Ch. 3 
In the previous chapter, we covered the basics of GO data types and, in a few examples, used flow control to print messages.
In this chapter, we'll use what we learned about comparison operators and boolean values to learn more about flow control and its patterns and build some games to show at your winter holiday parties.
Elements of Flow control
Programs are like water in a river. They tend to flow in one direction. And like water in a river, parts of it can go off into separate outlets based upon the river's conditions, such as a boulder in the stream or a dam.
However, unlike water in a river, it's much easier to divert the flow of a program.
To do so, we use conditions to evaluate a variable's value or the output of a function and execute a different code block.
Let's use an example with comments to see.
We are using the familiar 'if' from the previous chapter.

In this program, we're checking whether a 'random' integer number is even or not via the modulo operator %, which is a fancy way of saying 'dividing with remainders.
If you recall, an even number can be divided in two without any remainders, whereas an odd number cannot.
Don't worry about the math. Just note that the logic of the program changes based on the condition, which is the 'random' value of the n  variable.
Now, let's introduce the rest of the gang for flow control.
Flow Control Statements
'if' statements
These are the most used control statements and the safe go-to for starting. The general pattern is that your program has a chunk of code that will only be executed if the condition is considered 'true' or 'false.' As shown above.
'else' statements
Technically, in the previous 'if' example, we could've had an 'else' statement to print the 'number is odd:' part, but we kept it cleaner by leaving it out (it was implicitly handled), which is a preference amongst some programmers.  However, as a newcomer, it's better to do more showing:

Else statements say, 'hey, if this piece of logic turns out to be false or something else, then do this.'
Typically, 'else' statements aren't preferred because they aren't as explicit and will catch only when the condition isn't what is expected.
Let's introduce the else's cousin.
else if
Here, we can communicate what we want to happen 'if' the first condition evaluates to 'false' and explicitly define what should happen.

switch
The 'switch' control flow works much the same as using 'if' 'else', or 'else if' but presents differently in the form of 'cases'. A 'case' is the same as a condition and may be preferred by those with a mathematical background. It is, in this Author's opinion, the most explicit way to share what the code is training to do.

In the above, we say 'switch' based upon the following 'case' and pass in some data. The same again can be accomplished by using 'if' and 'else if'.


In general, I prefer using 'switch' and including the optional 'default' case because I find it easier to test my 'cases' and to convey to others what I'm trying to do with my code.
Now that we've added a few more things to our tool belt let's add the final section for this chapter for making fun programs, also known as games.
For loops
Now that we have some logic to control a program's flow, it's time to introduce 'loops'. A loop is a way we tell computers to "do work for x amount of times".
A 'for' loop also uses a condition to do its work, but it works 'until' that condition is satisfied.
So, for example, imagine if we took the common Norwegian saying 'a thousand thanks' literally.

In the above, the 'for' loop syntax is for initial condition; compare;if not met, then do another { // work to be done }
In most programming languages, it is common to start at '0', so keep this in mind when working with arrays, slices, and initial loop conditions.
Games
Okay, enough theory and toy examples; let's have fun by building a few small games to show your friends and family at the winter holiday parties.
Don't worry if some syntax is unfamiliar. We want to focus on building muscle memory and familiarity with GO. So, try your best to copy the code and get in running. If you have any troubles at all, the source code is available
Guess the Number
In this first game, we'll have the computer pick a random number and then try to guess it by inputting guesses, but the catch is we only get '3' attempts!
But, we'll give the user clues to increase the chances of winning within 3.

As I hope you can see, using the 'switch' statement clarifies what 'cases' we are evaluating for.  If you want an additional exercise, try writing this with 'if' 'else' and 'else if' statements.  
Advanced: Rock, Paper, Scissors
Now, for a more fun and advanced example. Let's whet your programmer appetite by seeing what fun can be had with loops and control flows.
In this code, you'll see advanced GO things like 'type' and 'maps', which we'll spend a whole section on in 'Data structures'  later.
But for now, just code out what you see to get familiar with the syntax.

Above, you encountered a lot of new things but saw one familiar concept in the 'for' loop. You'll notice there are no conditions with it, so it will keep running forever. It's pretty neat because a forever-running loop is a basis for making a videogame.