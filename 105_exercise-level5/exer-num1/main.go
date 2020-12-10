package main

import "fmt"

type person struct {
	fname  string
	lname  string
	flavor []string
}

func main() {
	p1 := person{
		fname: "James",
		lname: "Bond",
		flavor: []string{
			"Pandan",
			"Ube",
		},
	}

	p2 := person{
		fname: "Marvin",
		lname: "Mahilum",
		flavor: []string{
			"Chocolate",
			"Mango",
		},
	}

	fmt.Println(p1.fname)
	fmt.Println(p1.lname)

	for i, v := range p1.flavor {
		fmt.Println("\t", i, v)
	}

	fmt.Println(p2.fname)
	fmt.Println(p2.lname)

	for i, v := range p2.flavor {
		fmt.Println("\t", i, v)
	}
}
