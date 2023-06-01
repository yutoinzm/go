package main

import (
	"fmt"
)

func main() {
	i := 0
x:
	for i < 5 {
		for k := 0; k < 5; k++ {
			if k == i {
				i++
				fmt.Printf("k and i is %d", k)
				goto x
			}
			fmt.Printf("  *\n")
		}
	}
}
