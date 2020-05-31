//Echo3 prints its command-line arguments using string.Join: 'echo1 "arg1" arg2 ...'
package main

import (
	"fmt"
	"os"
	"strings"
	"time" //Exercise 1.3, compare times of execution between echo1/2/3
)

func main() {
	start := time.Now()     //Exercise 1.3, compare times of execution between echo1/2/3
	fmt.Println(os.Args[0]) // Exercise 1.1, print the name of the function.
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.5fs elapsed\n", time.Since(start).Seconds()) //Exercise 1.3, compare times of execution between echo1/2/3
}

//Notes
/*
fmt.Println(os.Args[1:]) will print the same as strings but with surrounding brackets.
strings.Join is simpler and more efficient that previous loops.

Exercise 1.3 using 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 as test
echo1 0.0100s
echo2 0.0150s
echo3 <0.0010s
*/
