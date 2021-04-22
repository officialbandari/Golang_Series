/*
control statements of go

if
if else
else if
nested if


LOGICAL AND - &
LOGICAL OR - ||
LOGICAL NOT - !


*/

package main

import "fmt"

func main() {
	// var name string
	// var age int
	// fmt.Scan(&name)
	// fmt.Scan(&age)

	// if condtion

	// if name == "krishna" {
	// 	fmt.Println("welcome to officialTechs")

	// }

	//if else condition

	// if name == "krishna" {
	// 	fmt.Println("welcome to officialTechs")

	// } else {
	// 	fmt.Println("please contact to this no:7777")
	// }
	//voting example

	// if age >= 18 && age <= 100 {

	// 	fmt.Println(name, "ur eligible for voting")

	// } else {
	// 	fmt.Println(name, "your not eligible for voting")
	// }

	var val1, val2, val3 int
	fmt.Scan(&val1, &val2, &val3)

	if val1 > val2 && val1 > val3 {
		fmt.Println(val1, "is the greatest..")
	} else if val2 > val3 {
		fmt.Println(val2, "is the greates")
	} else {
		fmt.Println(val3, "is the greatest...")
	}

	fmt.Println("The program Ended...")

}
