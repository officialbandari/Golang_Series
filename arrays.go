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
	farr[0] = 1.1
	farr[1] = 2.0
	farr[2] = 3.1
	farr[3] = 4.5

	fmt.Println(farr)

	// string array
	var sarr [5]string

	sarr[0] = "krishna"
	sarr[1] = "bandari"
	fmt.Printf("%q \n", sarr)

	var barr [5]bool
	barr[0] = true
	barr[1] = false
	barr[2] = false
	barr[3] = true

	fmt.Println(barr)

	fmt.Println("end of the array...")
}
