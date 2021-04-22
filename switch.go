//switcht statement

package main

import "fmt"

func main() {

	// var n int

	// fmt.Scan(&n)

	// switch n {

	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")

	// default:
	// 	fmt.Println("Wrong Option")

	// }

	// condtion using
	// var n int
	// fmt.Scan(&n)

	// switch {

	// case n > 7:
	// 	fmt.Println(n, "The value greter than 7")
	// case n < 7:
	// 	fmt.Println(n, "The value is less than 7")
	// case n == 7:
	// 	fmt.Println(n, "The value is equal to 7")
	// default:
	// 	fmt.Println("The wrong option...")

	// }

	//days condition
	//falthrough is special condition in go
	// falthrough is after condition it ll print all the values

	var ch int
	fmt.Scan(&ch)

	switch ch {
	case 1:
		fmt.Println("Monday")
		fallthrough
	case 2:
		fmt.Println("Tuesday")
		fallthrough
	case 3:
		fmt.Println("Wednesday")
		fallthrough
	case 4:
		fmt.Println("Thursday")
		fallthrough
	case 5:
		fmt.Println("Friday")
	}

}
