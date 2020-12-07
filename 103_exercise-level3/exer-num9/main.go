package main

import "fmt"

func main() {
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("Go to the mountains")
	case "swimming":
		fmt.Println("Go to the pool")
	case "surfing":
		fmt.Println("Go to the Hawaii")
	case "wingsuit flying":
		fmt.Println("What would you like me to say to you")
	}
}
