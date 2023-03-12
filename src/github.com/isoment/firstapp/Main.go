package main

import "fmt"

/************
* VARIABLES *
************/

// // We can declare variables in blocks
// var (
// 	firstName  string = "Bill"
// 	lastName   string = "Smith"
// 	employeeId string = "ID4823"
// )

// func main() {
// 	fmt.Println("Variables")
// 	var i int = 42
// 	fmt.Printf("%v, %T\n", i, i)

// 	// Variables can be converted from one type to another. We need to be careful with
// 	// this since we can lose information during the conversion process. Ie going from
// 	// a float to a number we can lose precision.
// 	var j float32
// 	j = float32(i)
// 	fmt.Printf("%v, %T\n", j, j)

// 	// When converting strings we can use the "strconv package" without this package Go
// 	// will end up converting the int to it's corresponding unicode char.
// 	var k string
// 	k = strconv.Itoa(i)
// 	fmt.Printf("%v, %T\n", k, k)
// }

/*************
* Primitives *
*************/

func main() {
	fmt.Println("Primitives")
	var a bool = true
	fmt.Printf("%v, %T\n", a, a)

	// The default integer type is int
	b := 42
	fmt.Printf("%v, %T\n", b, b)

	// Arithmetic operations
	fmt.Println("Numeric Operations")
	c := 10
	fmt.Println(b + c)
	fmt.Println(b - c)
	fmt.Println(b * c)
	fmt.Println(b / c)
	fmt.Println(b % c)

	// Arithmetic operations
	fmt.Println("Bit Operations")
	d := 10 // 1010
	e := 3  // 0011
	/*
		The And operator will check what bits are set in the first number and the
		second number at the given position.

		1010
		0011
		----
		0010
	*/
	fmt.Println(d & e) // 0010 ie 2 in decimal
	/*
		The Or operator will check if one OR the other is set at a given position

		1010
		0011
		----
		1011
	*/
	fmt.Println(d | e) // 1011 ie 11 in decimal
	/*
		The XOR Exclusive Or checks if one has the bit set OR the other does but not both

		1010
		0011
		----
		1001
	*/
	fmt.Println(d ^ e) // 1001 ie 9 in decimal
	/*
		The bit will be 1 if neither of the numbers have the bit set at a given position

		1010
		0011
		----
		0100
	*/
	fmt.Println(d &^ e) // 0100 ie 4 in decimal

	fmt.Println("Bit Shifts")
	f := 60 // 00111100
	// Bitshift left will shift the bit sequence 3 places to the left
	fmt.Println(f << 3) // 111100000
	// Bitshift right will shift the bit sequence 3 places to the right
	fmt.Println(f >> 3) // 111

	fmt.Println("Complex numbers")
	var g complex64 = 1 + 2i
	var h complex64 = 2 + 5.2i
	fmt.Println(g + h)
	fmt.Println(g - h)
	fmt.Println(g * h)
	fmt.Println(g / h)

	fmt.Println("Strings")
	i := "this is a string"
	// returns 105 uint8
	fmt.Printf("%v, %T\n", i[2], i[2])
	// we can convert to the actual char
	fmt.Printf("%v, %T\n", string(i[2]), i[2])
	j := "this is random"
	jb := []byte(j)
	fmt.Printf("%v, %T\n", jb, jb)

	fmt.Println("Runes")
	// We declare a rune using single quotes.
	rune := 'a'
	// 97, int32 because a is decimal 97 in unicde and a rune type is just an alias for int32
	fmt.Printf("%v, %T\n", rune, rune)
}
