package main

import "fmt"

var str string = "String declared" // can also use back quotes like `String format`

var n = 423

func main() {

	fmt.Println("Hello Go. What is it about you? ")

	//declare variable using short declaration operator.
	x := 44
	// each line is a statement.
	fmt.Println(x)
	// operator and operands
	y := 99 + 24
	fmt.Println(y)
	fmt.Println("%T", str)

	// built in datatypes vs composite data type
	name := "Amogh"
	fmt.Print(name)
	fmt.Println("\n Well, I really am ", name, " !")

	fmt.Printf(name)
	fmt.Printf("%T\n", name)
	fmt.Printf("%b\n", n) // binary
	fmt.Printf("%x\n", n) // hex
	fmt.Printf("%x\n", n) // hex

}
