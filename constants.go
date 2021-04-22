package main

import "fmt"

func main() {

	var val1 int = 32
	val1++

	const val2 = 33
	// we can't increase or decrease the constants value it prints the as default

	fmt.Println(val1, val2)
	//only constants value declaration

	const (
		num1 int   = 33
		num2 int32 = 44
		num3 int64 = 55
	)

	fmt.Println(num1, num2, num3)

	//and also we have special constant variable value that is iota

	/*

		we can't declare the value as below
		const num := somevalue
		const num := 3
		it will throw the error


	*/

	const (
		num4 = iota //0 is the default starting value
		num5 = iota //1
		num6 = iota //2
	)

	fmt.Println(num4, num5, num6)

}
