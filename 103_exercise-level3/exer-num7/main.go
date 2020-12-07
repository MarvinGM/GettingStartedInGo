package main

import "fmt"

func main() {
	x := "James Bond"
	if x == "Hehe" {
		fmt.Println("Nope")
	} else if x == "James Bond" {
		fmt.Println("Yes it is!")
	} else {
		fmt.Println("Neither")
	}
}
