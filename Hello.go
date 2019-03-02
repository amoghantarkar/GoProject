package main

import "fmt"

// if you need to declare variable outside func body then use var.
var z = 90

// no initialization then it is 0 for int and boolean - false, "" for string, nil for pointers, functions, interfaces,
//slices, channels, maps

var m int

func main() {
	fmt.Println("Hello Go. What is it about you? ")

	n, err := fmt.Print("Hello Again")

	fmt.Println(n)
	fmt.Println(err)
	//declare variable using short declaration operator.
	x := 44
	// each line is a statement.
	fmt.Println(x)
	// operator and operands
	y := 99 + 24
	fmt.Println(y)

	fmt.Println(z)

	test()
}

func test() {
	fmt.Print("inside test func")
}
