//Echo2 prints its command-line arguments using a range loop: 'echo1 "arg1" arg2 ...'
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""        //short variable declaration
	fmt.Println(os.Args[0]) // Exercise 1.1, print the name of the function.
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

//Notes
/*
In each iteration, range produces a pair of values: the index and the value of the element at that index.
The syntax of a range loop requires inclusion of the index.
Since an unused variable declaration is illegal, we use the _ or blank identifier.
The blank indentifier may be used whenever syntax requires a variable name but program logic does not.
These declarations:
	1. s := ""
	2. var s string
	3. var s = ""
	4. var s string = ""
are all valid and equivalent.
The short declaration (1) may only be used within a function, not for package-level variables.
The second form (2) relies on default initialization to zero value.
The third form (3) is rarely used except when declaring multiple variables.
The fourth form (4) is explicit about the variable's type, which is redeundant when it is the same as that of the initial value, but necessary in other cases where they are not of the same type.
In practice, you should generally use one of the first two forms, with explicit initialization to say that the initial value is important and implicit initialization to say that the initial value doesn't matter.
On each iteration, s gets a new value. The old value is garbage-collected in due course.
*/
