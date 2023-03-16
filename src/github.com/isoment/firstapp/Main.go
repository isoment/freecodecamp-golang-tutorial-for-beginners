package main

import (
	"fmt"
	"math"
)

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

// func main() {
// 	fmt.Println("Primitives")
// 	var a bool = true
// 	fmt.Printf("%v, %T\n", a, a)

// 	// The default integer type is int
// 	b := 42
// 	fmt.Printf("%v, %T\n", b, b)

// 	// Arithmetic operations
// 	fmt.Println("Numeric Operations")
// 	c := 10
// 	fmt.Println(b + c)
// 	fmt.Println(b - c)
// 	fmt.Println(b * c)
// 	fmt.Println(b / c)
// 	fmt.Println(b % c)

// 	// Arithmetic operations
// 	fmt.Println("Bit Operations")
// 	d := 10 // 1010
// 	e := 3  // 0011
// 	/*
// 		The And operator will check what bits are set in the first number and the
// 		second number at the given position.

// 		1010
// 		0011
// 		----
// 		0010
// 	*/
// 	fmt.Println(d & e) // 0010 ie 2 in decimal
// 	/*
// 		The Or operator will check if one OR the other is set at a given position

// 		1010
// 		0011
// 		----
// 		1011
// 	*/
// 	fmt.Println(d | e) // 1011 ie 11 in decimal
// 	/*
// 		The XOR Exclusive Or checks if one has the bit set OR the other does but not both

// 		1010
// 		0011
// 		----
// 		1001
// 	*/
// 	fmt.Println(d ^ e) // 1001 ie 9 in decimal
// 	/*
// 		The bit will be 1 if neither of the numbers have the bit set at a given position

// 		1010
// 		0011
// 		----
// 		0100
// 	*/
// 	fmt.Println(d &^ e) // 0100 ie 4 in decimal

// 	fmt.Println("Bit Shifts")
// 	f := 60 // 00111100
// 	// Bitshift left will shift the bit sequence 3 places to the left
// 	fmt.Println(f << 3) // 111100000
// 	// Bitshift right will shift the bit sequence 3 places to the right
// 	fmt.Println(f >> 3) // 111

// 	fmt.Println("Complex numbers")
// 	var g complex64 = 1 + 2i
// 	var h complex64 = 2 + 5.2i
// 	fmt.Println(g + h)
// 	fmt.Println(g - h)
// 	fmt.Println(g * h)
// 	fmt.Println(g / h)

// 	fmt.Println("Strings")
// 	i := "this is a string"
// 	// returns 105 uint8
// 	fmt.Printf("%v, %T\n", i[2], i[2])
// 	// we can convert to the actual char
// 	fmt.Printf("%v, %T\n", string(i[2]), i[2])
// 	j := "this is random"
// 	jb := []byte(j)
// 	fmt.Printf("%v, %T\n", jb, jb)

// 	fmt.Println("Runes")
// 	// We declare a rune using single quotes.
// 	rune := 'a'
// 	// 97, int32 because a is decimal 97 in unicde and a rune type is just an alias for int32
// 	fmt.Printf("%v, %T\n", rune, rune)
// }

/************
* Constants *
************/

// // iota is a counter we can use when creating enumerated constants. Iota is scoped to a constant
// // block so if we declare another constant block itoa will be reset to 0 for that block. It is a
// // good idea to use the first as an error in case we forget to declare a variable value later and
// // compare it to the iota
// const (
// 	error = iota // 0
// 	cat          // 1
// 	dog          // 2
// 	snake        // 3
// )

// // Using a const block and iota to calculate how many bytes
// const (
// 	_  = iota             // ignore first value by assigning to blank identifier
// 	KB = 1 << (10 * iota) // 1024				iota is 1
// 	MB                    // 1048576			iota is 2 etc...
// 	GB                    // 1073741824
// 	TB                    // 1099511627776
// )

// /*
// Using const to set some role permission values. Each constant will represent one location in a byte,
// the first number is the decimal and the second the binary representation. We are just shifting to
// the left one binary place.

// We are able to use one byte to store 8 different permission roles. For example...

// 100101 has the roles isAdmin, canSeeFinancials and canSeeEurope
// */
// const (
// 	isAdmin            = 1 << iota // 1,    1
// 	isHeadquarters                 // 2,    10
// 	canSeeFinancials               // 4		100
// 	canSeeAfrica                   // 8		1000
// 	canSeeAsia                     // 16	10000
// 	canSeeEurope                   // 32	100000
// 	canSeeNorthAmerica             // 64	1000000
// 	canSeeSouthAmerica             // 128	10000000
// )

// func main() {
// 	fmt.Println("Constants")
// 	const myConst float64 = 1.5453454
// 	fmt.Printf("%v, %T\n", myConst, myConst)

// 	// Notice in this case how this is allowed. Even though the const is inferred to be of int
// 	// type it is still allowed to be added to an int16 type. This works because the compiler
// 	// is replacing const a with a literal value of the constant.
// 	const a = 10
// 	var b int16 = 2
// 	fmt.Printf("%v, %T\n", a+b, a+b)

// 	var animal int = dog
// 	fmt.Printf("%v\n", animal == dog)

// 	// Calculating filesize
// 	fileSize := 4000000000
// 	sizeGB := float32(fileSize) / GB
// 	fmt.Printf("%.2fGB\n", sizeGB)

// 	// Using the roles const from above, when we or the above values together we get a
// 	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
// 	fmt.Printf("%b\n", roles)
// 	/*
// 		000001 isAdmin
// 		100101 roles
// 		------
// 		000001

// 		The expression below evaluates to true
// 	*/
// 	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
// }

/*********
* Arrays *
*********/

// func main() {
// 	fmt.Println("Arrays")

// 	// Declaring the array and preassigning it values
// 	grades := [3]int{90, 100, 85}
// 	fmt.Printf("Grades: %v\n", grades)

// 	// Declaring the array and assigning values after the fact
// 	var students [3]string
// 	students[0] = "Lisa"
// 	students[1] = "Bill"
// 	students[2] = "Frank"
// 	fmt.Printf("Students: %v\n", students)
// 	fmt.Printf("Student Count: %v\n", len(students))

// 	// Arrays can be made up of arrays, ie a matrix
// 	var identityMatrix [3][3]int
// 	identityMatrix[0] = [3]int{1, 0, 0}
// 	identityMatrix[1] = [3]int{0, 1, 0}
// 	identityMatrix[2] = [3]int{0, 0, 1}
// 	fmt.Printf("Identity matrix: %v\n", identityMatrix)

// 	a := [...]int{1, 2, 3}
// 	b := a
// 	b[1] = 5
// 	fmt.Println(a)
// 	// We can see that b is a new copy of array a and it is not simply pointing to it in memory.
// 	fmt.Println(b)
// 	// We can use a pointer, which is just a reference to the original value in memory. Here we set
// 	// the value of c to point to a
// 	c := &a
// 	fmt.Println(c)
// 	// When we change an element in c it is reflected in a since c is just a pointer to a
// 	c[1] = 16
// 	fmt.Println(c)
// 	fmt.Println(a)
// }

/*********
* Slices *
*********/

// func main() {
// 	fmt.Println("Slices")

// 	// Declaring a slice and initializing some data
// 	a := []int{90, 100, 85}
// 	fmt.Println(a)

// 	// Just like with arrays we can call len() to get the length
// 	fmt.Printf("%v\n", len(a))

// 	// There is a function called cap() which gives the value of the underlying array
// 	fmt.Printf("%v\n", cap(a))

// 	// We can also use this syntax to create slices...
// 	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	// For the [x:x] syntax the first number is inclusive (included) and the second is exclusive (not included)
// 	c := b[:]   // Create a slice of all the elements in a
// 	d := b[3:]  // Create a slice from the 4th element to the end
// 	e := b[:6]  // Slice the first 6 elements
// 	f := b[3:6] // Slice the 4th, 5th and 6th elements
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// 	fmt.Println(e)
// 	fmt.Println(f)

// 	// We can use the make() function to create a slice, the first argument is the slice and type, the second
// 	// is the length and the third is the capacity.
// 	g := make([]int, 3, 100)
// 	fmt.Println(g)
// 	fmt.Printf("%v\n", len(g))
// 	fmt.Printf("%v\n", cap(g))

// 	h := []int{}
// 	fmt.Println(h)
// 	fmt.Printf("Length: %v\n", len(h))
// 	fmt.Printf("Capacity: %v\n", cap(h))
// 	h = append(h, 1, 2, 3, 4, 5, 6)
// 	fmt.Printf("Length: %v\n", len(h))
// 	fmt.Printf("Capacity: %v\n", cap(h))

// 	// Go has an operator similar to the spread operator that we can use to concat slices
// 	i := []int{1}
// 	i = append(i, []int{2, 3, 4, 5}...)
// 	fmt.Println(i)

// 	// We can remove the first element of a slice...
// 	j := []int{1, 2, 3, 4, 5}
// 	k := j[1:]
// 	fmt.Println(k)

// 	// We can remove the last element of a slice...
// 	l := []int{1, 2, 3, 4, 5}
// 	m := j[:len(l)-1]
// 	fmt.Println(m)

// 	// Removing elements from other positions is more complex, we have to create two slices and concatenate
// 	// them together, for example to remove the element at index 2. We need to be very careful with this since
// 	// the underlying array is being modified. If we are using the original slice anywhere else this will
// 	// cause problems.
// 	n := []int{1, 2, 3, 4, 5}
// 	o := append(n[:2], n[3:]...)
// 	fmt.Println(o)
// }

/*******
* Maps *
*******/

// func main() {
// 	townPopulations := map[string]int{
// 		"Faketown":  7434,
// 		"Fakeville": 420012,
// 		"Fakeburg":  472384,
// 		"Fakeglenn": 9182,
// 	}

// 	// Getting a value for a key
// 	fmt.Println(townPopulations["Fakeglenn"])
// 	// Setting a new entry in the map
// 	townPopulations["Realtown"] = 9236621
// 	// deleting a value from the map
// 	delete(townPopulations, "Realtown")
// 	fmt.Println(townPopulations)
// 	// We can use this syntax to return a boolean if the key is in the map or not
// 	pop, ok := townPopulations["Fakeburg"]
// 	fmt.Println(pop, ok)
// 	// We can get the count of elements in a map
// 	fmt.Println(len(townPopulations))
// }

/**********
* Structs *
**********/

// type Animal struct {
// 	Legs       int
// 	Name       string
// 	Species    string
// 	Attributes []string
// }

// /*
// Go supports embedding structs within other structs. Go favors composition and delegates
// the functionality to the Animal struct. We can access Animal values just like if they were part
// of the Bird struct. Bird is not an instance of an Animal it HAS A Animal
// */
// type Bird struct {
// 	Animal
// 	Speed  float32
// 	CanFly bool
// }

// type Request struct {
// 	Name   string `required max:"100"`
// 	Origin string
// }

// func main() {
// 	dog := Animal{
// 		Legs:       4,
// 		Name:       "Rover",
// 		Species:    "Canis Familiaris",
// 		Attributes: []string{"Furry", "Smelly", "Likes to drool"},
// 	}
// 	fmt.Println(dog)
// 	fmt.Println(dog.Attributes[0])

// 	// We can declare anonymous structs. The first {} defines the structure and the second
// 	// initializes the values.
// 	aDoctor := struct{ name string }{name: "Billford"}
// 	fmt.Println(aDoctor.name)
// 	// Creating a struct from another struct and then modifying the new one will not change the original
// 	anotherDoctor := aDoctor
// 	anotherDoctor.name = "Sally"
// 	fmt.Println(aDoctor.name)

// 	// We can declare a struct this way...
// 	bird := Bird{}
// 	bird.Name = "Emu"
// 	bird.Legs = 2
// 	bird.Speed = 48
// 	bird.CanFly = false
// 	fmt.Println(bird.Name)

// 	// When using the literal syntax to declare a struct we need to be more specific...
// 	hawk := Bird{
// 		Animal: Animal{Name: "Hawk", Legs: 2},
// 		Speed:  60,
// 		CanFly: true,
// 	}
// 	fmt.Printf("%v\n", hawk)

// 	// We can access a structs tags using reflection
// 	t := reflect.TypeOf(Request{})
// 	field, _ := t.FieldByName("Name")
// 	fmt.Println(field.Tag)
// }

/****************
* If Statements *
****************/

func main() {
	// Variables can be initialized in an if statement. We have access to the variables
	// initialized within the block scope. They are not available outside the block!
	a := map[string]string{
		"Faketown": "A great place to live",
	}
	if pop, ok := a["Faketown"]; ok {
		fmt.Println(pop)
	}

	// Go has a feature called short circuiting. We might think that the if statement below
	// would result in fmt.Println("returning true...") but it does not. This is because the first
	// part of the if statement guess < 1  results in true and they other two parts do not even get
	// evaluated.
	guess := -30
	if guess < 1 || returnTrue() || guess > 100 {
		fmt.Println("The guess must be between 1 and 100")
	}

	// We need to exercise caution when using logic operations an floating point numbers. The expected
	// behavior is that num an a are always equal
	num := 0.123
	if a := math.Pow(math.Sqrt(num), 2); a == num {
		fmt.Println(a)
		fmt.Println("These are the same")
	} else {
		fmt.Println(a)
		fmt.Println("These are different")
	}
}

func returnTrue() bool {
	fmt.Println("returning true...")
	return true
}
