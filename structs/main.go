package main

import "fmt"

type contactInfo struct {
	email   string
	pincode int
}

type person struct {
	firstName string
	lastName  string
	age       int
	contactInfo
}

func main() {
	var alex person
	alex.print()
	alex.firstName = "alex"
	alex.lastName = "anderson"
	alex.age = 30
	alex.contactInfo.email = "alex@mail.com"
	alex.contactInfo.pincode = 14120
	fmt.Println(alex)
	sam := person{"Sam", "Dane", 21, contactInfo{"sam@mail.com", 14001}}
	fmt.Println(sam)
	tany := person{
		firstName: "Tanny",
		lastName:  "Tense",
		age:       24,
		contactInfo: contactInfo{
			email:   "tanny@mail.com",
			pincode: 13001,
		},
	}
	fmt.Println(tany)
	tany.updateFirstName("Tany")
	tany.print()
}

func (pointerToPerson *person) updateFirstName(nameToUpdate string) {
	(*pointerToPerson).firstName = nameToUpdate
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
