package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	a := runtime.GOMAXPROCS(2)
	go helloName("ehsan")
	fmt.Println("Hello World")
	time.Sleep(1 * time.Second)
	fmt.Println(a)
	//lion := new(port.Lion)
	//lion.Name = "Lion"
	//lion.Age = 12
	//
	//cat := port.Cat{
	//	Name: "Cat",
	//}
	//fmt.Println(lion)
	//fmt.Println(&cat)
	//fmt.Println(GetNameAnimal(lion))
}
func helloName(name string) {
	fmt.Println("Hello " + name)
}

//func GetNameAnimal(animal port.Animal) string {
//	return animal.GetName()
//}
