package main

import "fmt"

var x = 42

func main() {
	fmt.Printf("Decimal: %d\n", x)
	fmt.Printf("Binary:\t%b\n", x)
	fmt.Printf("Hexa:\t%#x\n", x)
	y := x << 1
	fmt.Printf("Decimal: %d\n", y)
	fmt.Printf("Binary:\t%b\n", y)
	fmt.Printf("Hexa:\t%#x\n", y)

}
