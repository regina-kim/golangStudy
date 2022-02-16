package main

import "fmt"

type Human struct {
	Name  string
	Age   int
	Phone string
}

func printHuman(a Human) {
	fmt.Println("Name: ", a.Name)
	fmt.Println("Age: ", a.Age)
	fmt.Println("Phone: ", a.Phone)
}

func main() {
	var paul Human
	paul.Age = 20
	paul.Name = "Paul"
	paul.Phone = "010-0000-0000"

	printHuman(paul)

	var alice Human = Human{Name: "Alice", Phone: "010-1234-1234"}
	printHuman(alice)

	tom := paul
	tom.Name = "Tom"
	printHuman(tom)

}
