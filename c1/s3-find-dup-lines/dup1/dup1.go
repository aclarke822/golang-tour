//Dup1 prints the text of each line that appears more than once
//once in the standard input, preceded by its count.package lines
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	//NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//Notes
/*
As with for, parentheses are never used around the condition in an if statement, but braces are required for the body.
There can be an optional else that is executed if the condition is false.
*/
