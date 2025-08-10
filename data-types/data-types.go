package main

import "fmt"

func DataTypesInt() {
	// Example of integer data type signed ints can be negative or positive
	var a int8 = 10
	var b int16 = 20
	var c int32 = 30
	var d int64 = 40

	fmt.Println("Integer values:", a, b, c, d)
	// Unsigned integers cannot be negative
	var e uint8 = 50
	var f uint16 = 60
	var g uint32 = 70
	var h uint64 = 80
	fmt.Print("Unsigned integer values:", e, f, g, h)
}

func DataTypesFloat() {
	// Example of float data type
	var a float32 = 3.14
	var b float64 = 2.71828

	fmt.Println("Float values:", a, b)
}

func DataTypesString() {
	// Example of string data type
	var str1 string = "User"
	var str2 string = "World"

	fmt.Printf("Hello %s Welcome to my %s\n", str1, str2)

	// String Literal
	var str3 string = `
	This is a 
	multiline string literal 
	that preserves line breaks and spaces. 
	`
	fmt.Printf("%s\n", str3)
}
