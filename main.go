package main

import "fmt"

func main() {
	// a := infrastructure.NewSqlHandler()
	// cmd := `CREATE TABLE IF NOT EXISTS persons(
	// 			name INT,
	// 			age  INT)`
	// _, err := a.Conn.Exec(cmd)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	bob := new(Person)
	mike := new(Person2)

	IntroForPerson(bob)  //=> Hello World
	IntroForPerson(mike) //=> Hello World
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
