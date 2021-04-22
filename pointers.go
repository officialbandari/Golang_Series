// POINTER IN GO

package main

import "fmt"

func main() {

	// var num1, num2 int
	// var ptr *int // Pointer variable
	// & is a address of variable
	//* is value of the address variable
	// fmt.Scan(&num1)
	// num2 = num1

	// num2++
	// ptr = &num2

	// fmt.Println(&num1, &num2, "The values of", num1, num2, ptr, *ptr)

	// dereferncing value

	var val1 int
	var ptr *int // Pointer variable
	fmt.Scan(&val1)

	ptr = &val1

	*ptr = 33

	fmt.Println(ptr, *ptr)

}
