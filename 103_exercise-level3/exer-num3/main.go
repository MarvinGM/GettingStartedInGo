package main

import "fmt"

func main() {
	bd := 1997
	counter := 0

	for bd < 2020 {
		fmt.Println(bd)
		bd++
		counter++
	}

	fmt.Println("My age: ", counter)
}
