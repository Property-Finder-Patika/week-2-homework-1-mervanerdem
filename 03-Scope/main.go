//An exercise for iota
package main

import "fmt"

var x int64 = 10

func main() {
	fmt.Printf("From example file Pi:%v\n", Pi)
	fmt.Printf("Global x:%v , %T\n", x, x)

	x := float64(5)
	fmt.Printf("Main x:%v , %T\n", x, x)

	if true {
		fmt.Printf("Main x in if block:%v , %T\n", x, x)
		x := "Hi"
		fmt.Printf("If block x:%v , %T\n", x, x)
		fmt.Printf("From example file in if block Pi:%v\n", Pi)
	}

}
