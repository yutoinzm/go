package main

import (
	"fmt"
)

func main() {
	fmt.Println("counting.....")

	i := 0
	for i < 10 {
		defer fmt.Println(i)
		i++
	}

	fmt.Println("done")
}
