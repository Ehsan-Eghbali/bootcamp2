package bootstrap

import (
	"bootcamp2/repo"
	"fmt"
)

func InitApplication() {
	person := repo.Person{}
	person.SetName("John Doe")
	//person.GetType()
	fmt.Println(person.GetName())

	owl := new(repo.Animal)
	owl.SetType("owl")
	fmt.Println(owl.GetType())
}
