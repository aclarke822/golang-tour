//Echo2 prints its command-line arguments using a range loop: 'echo1 "arg1" arg2 ...'
package main

import (
	"fmt"
	"os"
	"strconv" //Exercise 1.2, print index and value, one per line
	"time"    //Exercise 1.3, compare times of execution between echo1/2/3, using 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20
)

func main() {
	start := time.Now()
	fmt.Println(os.Args[0]) // Exercise 1.1, print the name of the function.
	for i, arg := range os.Args[1:] {
		fmt.Println(strconv.Itoa(i+1) + ":" + arg) //Exercise 1.2, print index and value, one per line
	}
	fmt.Printf("%.4fs elapsed\n", time.Since(start).Seconds())
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

Exercise 1.3 using 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 as test
echo1 0.0100s
echo2 0.0150s
echo3 <0.0010s
*/
