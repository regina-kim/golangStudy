package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scan(&a, &b)
	fmt.Println(Add(a, b))
	// val, tf := Divide(a, b)
	if val, tf := Divide(a, b); tf {
		fmt.Println(val)
	} else {
		fmt.Println("can not divide with 0")
	}
}

func Add(a int, b int) int {
	return a + b
}

func Divide(a int, b int) (int, bool) {
	if b == 0 {
		return 0, false
	} else {
		return a / b, true
	}
}
