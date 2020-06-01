//Dup3 prints the count and text of each line that
//appears more than once in the named input files
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//Notes
/*
ReadFile returns a byte slice that must be converted into a string so it can be split by strings.Split.
Under the hood, bufio.Scanner, ioutil.Readfile and ioutil.Writefile use the Read and Write methods of *os.File.
*/
