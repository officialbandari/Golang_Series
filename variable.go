package main

import "fmt"

func main() {

	//operators

	fmt.Println(1 + 2)
	fmt.Println(3 - 1)
	fmt.Println(4 * 2)
	fmt.Println(10 / 2)
	fmt.Println(12 % 2)

	//variable declaration with var key word

	var name string = "krishna"
	var age int32 = 29
	var address string = "kakatiyanagar ashoknagar sangareddy"
	var pin int64 = 502032

	fmt.Println("my", name, "am", age, "The", address, "of am and ", pin)

	//another way of using var key word

	var (
		fname    string = "krishna"
		fage     int32  = 30
		faddress string = "kakatiyanagar"
		fpin     int64  = 502032
	)

	fmt.Println("The details", fname, fage, faddress, fpin)

	//The shortHand diclaration

	sname := "krishna"
	sage := 29
	saddress := "kakatiyanagar ashoknagarBHEL"
	spin := 502032

	fmt.Println(sname, sage, saddress, spin)

}
