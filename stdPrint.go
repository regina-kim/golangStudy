package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a int = 10
	var b int = 20
	f := 327.82
	tf := false

	fmt.Print("a: ", a, "b: ", b)
	fmt.Println("a: ", a, "b: ", b)
	fmt.Printf("a: %d b: %d f(int): %d f(float64): %f\n", a, b, int(f), f)
	fmt.Printf("tf: %t, type of tf: %s\n", tf, reflect.TypeOf(tf))
}
