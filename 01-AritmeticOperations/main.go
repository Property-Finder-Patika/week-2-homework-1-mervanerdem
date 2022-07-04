//An exercise for arithmetic operations
package main

import "fmt"

func add(x, y float64) float64 { // add 2 number
	return x + y
}

func sub(x, y float64) float64 { // subtract 2 number
	return x - y
}

func div(x, y float64) float64 { // divide 2 number
	if y == 0 {
		return 0
	}
	return x / y
}

func multip(x, y float64) float64 { // multiply 2 number
	return x * y
}

func higher(x, y float64) string { // find higher one
	if x > y {
		return "higher"
	} else if y > x {
		return "lower"
	}
	return "equal" //no need another if
}

func main() {
	x := 5.5
	y := 2.3
	a := add(x, y)
	s := sub(x, y)
	d := div(x, y)
	m := multip(x, y)
	h := higher(x, y)
	fmt.Printf("%v + %v = %v\n", x, y, a) //print values to screen
	fmt.Printf("%v - %v = %v\n", x, y, s)
	fmt.Printf("%v / %v = %0.5v\n", x, y, d)
	fmt.Printf("%v * %v = %0.5v\n", x, y, m)
	fmt.Printf("%v is %v than %v\n", x, h, y)

}
