//An exercise for conversions among basic types
//An exercise for conversions among string and basic types using strconv package
//An exercise for iota
package main

import (
	"fmt"
	"strconv"
)

func stoI(x string) int { //function of convert string to integer
	i, _ := strconv.Atoi(x) // convert
	return i
}

func itoS(x int) string { //function of convert integer to string
	str := strconv.Itoa(x) //convert
	return str
}

func main() {
	f := 10.5
	i := int(f)           // float to int
	s := itoS(i)          // int to string
	i32 := int32(stoI(s)) // string to int and int to int32(it's dangerous don't do it at home)
	i64 := int64(i32)     // int32 to int64

	fmt.Printf("f= %v %T\n", f, f) // print screen
	fmt.Printf("i= %v %T\n", i, i)
	fmt.Printf("s= %v %T\n", s, s)
	fmt.Printf("i32= %v %T\n", i32, i32)
	fmt.Printf("i64= %v %T\n", i64, i64)

}
