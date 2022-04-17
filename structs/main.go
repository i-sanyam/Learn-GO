package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	var alex person
	fmt.Printf("%+v", alex)
	alex.firstName = "alex"
	alex.lastName = "anderson"
	alex.age = 30
	fmt.Println(alex)
	sam := person{"Sam", "Dane", 21}
	fmt.Println(sam)
	tany := person{firstName: "Tanny", lastName: "Tense", age: 24}
	fmt.Println(tany)
	tany.firstName = "Tany"
	fmt.Println(tany)
}
