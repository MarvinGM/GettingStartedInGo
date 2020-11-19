package main

import (
	"fmt"
	"strings"
)

func main() {

	//String declaration first then assigning a value
	var a1 string
	a1 = "hello"
	fmt.Println("1 ", a1)

	//String declaration in one line
	b1 := "hi"
	fmt.Println("2 ", b1)

	//String declaration checking if the first variable contain the string in the 2nd variable
	//will display a boolean true/false
	c1 := "hi hello"
	c2 := "hello"
	fmt.Println("3 ", strings.Contains(c1, c2))

	//String declaration ReplaceAll function
	d1 := "hi hello"
	fmt.Println("4 ", strings.ReplaceAll(d1, "h", "MM"))

	//String concatination c1 and c2
	fmt.Println("5 ", c1+c2)
	fmt.Println("6 ", c1+" "+c2)

}
