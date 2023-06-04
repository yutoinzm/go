package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Tarou", 35}
	z := Person{"Tom Howland", 26}
	fmt.Println(a, z)
}
