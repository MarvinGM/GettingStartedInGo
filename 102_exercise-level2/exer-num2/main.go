package main

import "fmt"

func main() {
	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)
	fmt.Println("a :", a)
	fmt.Println("b :", b)
	fmt.Println("c :", c)
	fmt.Println("d :", d)
	fmt.Println("e :", e)
	fmt.Println("f :", f)

}
