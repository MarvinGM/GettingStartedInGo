package main

import "fmt"

func swap(b1, b2 *int) {
	var temp int
	temp = *b2
	*b2 = *b1
	*b1 = temp
}

func main() {

	a1 := 2
	ptr := &a1
	//displaying the address of the variable
	fmt.Println("Address of a1: ", ptr)
	//displaying the value of the variable
	fmt.Println("Value of a1: ", *ptr)

	b1, b2 := 2, 3
	fmt.Println("Before swap: ", b1, b2)
	swap(&b1, &b2)
	fmt.Println("After swap: ", b1, b2)
}
