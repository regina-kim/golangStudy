package main

import "fmt"

const numArr int = 5

// 단, 변수명을 array 길이로 정의할 수는 없다.
// const는 매크로라서 사용 가능

func main() {

	var t [numArr]int = [numArr]int{0, 2, 4, 6, 8}
	for i := 0; i < numArr; i++ {
		fmt.Println(t[i])
	}

	days := [numArr]string{1: "monday", 3: "tuesday"}
	for i, v := range days {
		fmt.Println(i, v)
		// index와 value를 한번에 출력
		// i가 index, v가 value
	}

	x := [...]int{10, 20, 30, 40, 50}
	for i := 0; i < numArr; i++ {
		fmt.Println(x[i])
	}

	var y [numArr]int
	y = x
	for i := 0; i < numArr; i++ {
		fmt.Println(y[i])
	}

}
