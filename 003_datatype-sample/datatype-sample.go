package main

import "fmt"

func main() {

	//solo delaration
	var a1 int
	a1 = 2
	fmt.Print("Solo Declaration: ")
	fmt.Println(a1)

	//multiple declaration
	var (
		b1 = 2
		b2 = 3
	)
	fmt.Print("Multiple Declaration: ")
	fmt.Println(b1 + b2)

	//different data type declaration using casting
	var c1 int32 = 3
	var c2 int64 = 6
	fmt.Print("Casting Declaration: ")
	fmt.Println(int64(c1) + c2)

	//one line declaration
	d1 := 4
	d2 := 4
	fmt.Print("OneLine Declaration: ")
	fmt.Println(d1 + d2)
}
