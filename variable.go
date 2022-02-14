package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a int = 10
	var sent string = "sentence\n"
	var b int // type별 기본값 0 자동으로 할당
	var c = 4 // 타입을 생략하면 우변 값의 타입이 변수로 할당
	d := 5    // 선언 대입문 := 를 이용해 vare 키워드와 타입 생략
	e := 0.0  // 0.0으로 선언하면 float64로 할당됨
	fmt.Println(a, b, c, d, e, sent)
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(e))

	var tr bool = true
	fmt.Println(tr)

}
