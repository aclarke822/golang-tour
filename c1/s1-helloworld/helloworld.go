package main

import "fmt"

func main() {
	fmt.Println("Hello, 你好")
}

//Notes
/*
Run with go run helloworld.go
Build with go build helloworld.go

Go is organized into packages.
A package consists of one or more .go source files in a single directory that defines what a package does.
Each source file begins with a package declaration.
The Go standard library has over 100 packages.
The fmt package contains functions for printing formatted output and scanning input.
fmt.Println(string) outputs string to console with new line.

Package main defines a standalone executable program.
The function main is where program execution begins.
import declarations must follow the package declaration.
A program consists of one or more func, var, const and type.

A function declaration consists of the keyword func, the name of the function, a parameter list (main is empty), a result list (main is empty) and the body.
Go does not require semicolons at the ends of statements or declarations except where two or more appear on the same line.
The gofmt tool rewrites code into the standard format.

Use goimports to manage insertion and removal of import declarations.
$go get golang.org/x/tools/cmd/goimports

*/
