package main

import "fmt"

type Data struct {
	idx  int
	Data string
}

func main() {
	var p *int
	var a int
	p = &a

	a = 5
	fmt.Println(a, *p)

	*p = 20
	fmt.Println(a, *p)

	if p == nil {
		fmt.Println("pointer points an empty value")
	} else {
		fmt.Println("pointer points: ", *p)
	}

	var strPtr *Data = &Data{
		0, "data",
	}
	var strPtr2 *Data
	strPtr2 = strPtr

	fmt.Println(strPtr.idx, strPtr.Data)
	fmt.Println(strPtr2.idx, strPtr2.Data)
}
