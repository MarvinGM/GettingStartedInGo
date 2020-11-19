package main

import "fmt"

type egg int

var x egg

func main() {

	fmt.Println("Value of x: ", x)
	fmt.Printf("Type of x: %T\n", x)
	x = 42
	fmt.Println("New value of x: ", x)
}
