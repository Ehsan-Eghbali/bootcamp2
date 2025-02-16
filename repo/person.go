package repo

import "fmt"

type Person struct {
	name     string //""
	Family   string //""
	Age      int    //0
	IsActive bool   //false
}

//
//func NewPerson() *Person {
//	return &Person{}
//}

//func (r *Person) GetName() string {
//	return r.name
//}
//func (r *Person) SetName(name string) {
//	r.name = name
//}

func HelloWorld() {
	fmt.Println("Hello World")
}
