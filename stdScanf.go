package main

import "fmt"

func main() {

	var a1 int
	var b1 int
	n1, err1 := fmt.Scan(&a1, &b1)
	fmt.Println(n1, err1)
	// 정상적인 input이 들어오면 err는 <nil>, n은 총 input의 개수
	if err1 == nil {
		fmt.Println("sucessfully inserted")
	} else {
		fmt.Println("invalid input")
	}

	var a2 int
	var b2 int
	n2, err2 := fmt.Scanf("%d %d\n", &a2, &b2)
	fmt.Println(n2, err2)
	// 정상적인 input이 들어오면 err는 <nil>, n은 총 input의 개수
	if err2 == nil {
		fmt.Println("sucessfully inserted")
	} else {
		fmt.Println("invalid input")
	}

}
