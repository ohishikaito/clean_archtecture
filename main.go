package main

import (
	"fmt"
)

func main() {
	bob := new(Person)
	mike := new(Person2)

	IntroForPerson(bob)  //=> Hello World
	IntroForPerson(mike) //=> Hello World

	pochi := new(Dog)
	NakuForAnimail(pochi)
}

type Person struct{}  //Person構造体
type Person2 struct{} //Person2構造体

type People interface {
	intro()
}

func IntroForPerson(arg People) {
	arg.intro()
}

//Person構造体のメソッドintro()
func (p *Person) intro() {
	fmt.Println("Hello World")
}

//Person2構造体のメソッドintro()
func (p *Person2) intro() {
	fmt.Println("Hello World")
}

type Dog struct{}

func (d *Dog) naku() {
	fmt.Println("wan!")
}

type Animal interface {
	naku()
}

func NakuForAnimail(arg Animal) {
	// arg.naku()
}
