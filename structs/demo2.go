package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) Save() {
	fmt.Println("Eklendi: ", c.firstName)
}

func (c customer) Update() {
	fmt.Println("Güncellendi: ", c.firstName)
}

func Demo2() {
	c := customer{firstName: "Enes", lastName: "Gurel", age: 21}
	c.Save()
	c.Update()
}
