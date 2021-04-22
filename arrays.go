//ARRAYS IN GO

package main

import "fmt"

func main() {
	//integer array
	var arr [5]int

	arr[0] = 1
	arr[1] = 1
	arr[2] = 2
	fmt.Println(arr)

	//float array
	var farr [5]float32
	farr[0] = 1
	farr[1] = 2
	farr[2] = 3
	farr[3] = 4

	fmt.Println(farr)

	// string array
	var sarr [5]string
	fmt.Printf("%q \n", sarr)

	var barr [5]bool
	fmt.Println(barr)

	fmt.Println("end of the array...")
}
