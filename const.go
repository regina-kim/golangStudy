package main

import "fmt"

var a int

const constVal int = 10

const (
	Red   int = iota
	Blue  int = iota
	Green int = iota
)

func printColor(a int) {
	switch a {
	case Red:
		fmt.Println("Red")
	case Blue:
		fmt.Println("Blue")
	case Green:
		fmt.Println("Green")
	default:
		fmt.Println("Not a color")
	}
	// if a == Red {
	// 	fmt.Println("Red")
	// } else if a == Blue {
	// 	fmt.Println("Blue")
	// } else {
	// 	fmt.Println("Green")
	// }
	return
}

func checkColor(a int) {
	switch a {
	case Red, Blue, Green:
		fmt.Println("Is a Color")
	default:
		fmt.Println("Not a Color")
	}
}

func main() {
	a = 1
	// constVal = 20
	fmt.Println(a, constVal)
	fmt.Println(Red, Blue, Green)

	fmt.Scan(&a)
	// printColor(a)
	checkColor(a)
}
