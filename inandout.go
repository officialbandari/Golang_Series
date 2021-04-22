/*
// format specifier / Verb in Go

%v - value
%T - Type of the value
%d - Numeric data (int/uint) and integer type
%b - binary representation
%o - octal representation
%x /%X - hexa representation

%e - scientific natation
%f - decimal point value (float32, float64, complex64, complex128)
%g - decimal point values with neccessary digits (float32, float64, complex64, complex128)

%s - String

// 	Input Function
Scan, Scanln, Scanf --> standard input
Sscan, Sscanln , Sscanf --> buffer

//Output function

Print, Println, Printf --> standard function
Sprint, Sprintln, Sprintf -->buffer



*/

package main

import "fmt"

func main() {

	var name string
	var age int

	fmt.Scan(&name, &age)
	fmt.Printf("%s,%d", name, age)
}
