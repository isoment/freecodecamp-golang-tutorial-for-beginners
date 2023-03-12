package main

import (
	"fmt"
	"strconv"
)

// We can declare variables in blocks
var (
	firstName  string = "Bill"
	lastName   string = "Smith"
	employeeId string = "ID4823"
)

func main() {
	fmt.Println("Variables")
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	// Variables can be converted from one type to another. We need to be careful with
	// this since we can lose information during the conversion process. Ie going from
	// a float to a number we can lose precision.
	var j float32
	j = float32(i)
	fmt.Printf("%v, %T\n", j, j)

	// When converting strings we can use the "strconv package" without this package Go
	// will end up converting the int to it's corresponding unicode char.
	var k string
	k = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", k, k)

}
