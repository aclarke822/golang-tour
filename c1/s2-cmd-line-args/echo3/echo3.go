//Echo3 prints its command-line arguments using string.Join: 'echo1 "arg1" arg2 ...'
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args[0]) // Exercise 1.1, print the name of the function.
	fmt.Println(strings.Join(os.Args[1:], " "))
}

//Notes
/*
fmt.Println(os.Args[1:]) will print the same as strings but with surrounding brackets.
strings.Join is simpler and more efficient that previous loops.
*/
