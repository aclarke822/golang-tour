//Echo1 prints its command-line arguments: 'echo1 "arg1" arg2 ...'
package main

import (
	"fmt"
	"os"
	"strconv" //Exercise 1.2, print index and value, one per line
	"time"    //Exercise 1.3, compare times of execution between echo1/2/3
)

func main() {
	start := time.Now()
	fmt.Println(os.Args[0]) // Exercise 1.1, print the name of the function.
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(strconv.Itoa(i) + ":" + os.Args[i]) //Exercise 1.2, print index and value, one per line
	}
	fmt.Printf("%.4fs elapsed\n", time.Since(start).Seconds())
}

//Notes
/*
The os package provides functions and other values for dealing with the operating system in a platform independent fashion.
Command-line arguments are available to a program in a variable named Args as part of the os package, therefore it's name anywhere outside the os package is os.Args.
The variable os.Args is a slice of strings.
slice is a dynamically sized sequence s of array elemnts where individual elements can be accessed as s[i] and a contiguous subsequence as s[m:n].
The number of elements of slice is given as len(s).
Go uses half-open intervals.
The slice s[m:n], where 0<=m<=n<=len(s) contains n-m elements.
The first element of os.Args, os.Args[0] is the name of the command itself with the remainder being arguments presented at execution.

var declares 2 variables of type string on a single line, s and sep.
A variable can be initialized as part of its declaration, otherwise it is initialized to the zero value for its type (0 for numeric types and "" for strings).
The + operator concatenates strings.
The += assignment operator concatenates the original value to the left with the value on the right and assigns it to the var on the left.
Each logical operator like + or * has a corresponding assignment operator.
The := symbol is a part of a short variable declaration, a statement that declares one or more variables and gives them appropriate types based on the initializer values.
Increment/decrement statements (i++/i--) increment/decrement by 1 and are statements, not expressions. j = i++ is illegal.
The for loop is the only loop statement in Go with a number of forms:
	for initialization; condition; post {
		//zero or more statements
	}
Parantheses are never used around the components of a for loop. The braces are mandatory. The opening brace must be on the same line as the post statement (similar to func).
The initialization statement must be a simple statement, that is, a short variable declaration, an increment or assignment statement, or a function call.
The condition is a boolean expression that is evaluated at the beginning of each iteration to determine whether the statements in the loop are executed.
The post statement is executed after the body of the loop, then the condition is evaluated again.
The loop ends when the condition becomes false.
Any of these parts may be omitted. If there is no initialization and no post, the semicolons may be omitted.
A traditional 'while' loop:
	for condition {
		//zero or more statements
	}
A traditional infinite loop:
	for {
		//zero or more statements
	}

Exercise 1.3 using 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 as test
echo1 0.0100s
echo2 0.0150s
echo3 <0.0010s
*/
