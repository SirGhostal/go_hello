package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int in8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// Type does not need to be declared as it is inferred from the "Zach" string.
	var firstName = "Zach"
	var age int32 = 21
	var isStudent = true

	// Shorthand variable declaration
	lastName := "Howard"

	// Shorthand for multiple variables in one line
	email, eyeColour := "Zachary.Howard@outlook.com", "red"

	// You can print two types in one statement :O
	fmt.Println(firstName, lastName, age, isStudent, email, eyeColour)

	//https://godoc.org/fmt
	// %v	the value in a default format
	//	    when printing structs, the plus flag (%+v) adds field names
	// %#v	a Go-syntax representation of the value
	// %T	a Go-syntax representation of the type of the value
	// %%	a literal percent sign; consumes no value
	fmt.Printf("%T\n", age)
}