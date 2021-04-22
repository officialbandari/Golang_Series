//loops in golang

package main

import "fmt"

func main() {

	// var n int
	// fmt.Scan(&n)

	//sum := 0

	// for i := 0; i <= n; i++ {

	// 	sum = sum + i
	// }

	// fmt.Println(sum)

	// sum := 0

	// i := 0

	// for i <= n {
	// 	sum = sum + i
	// 	i++
	// }

	// fmt.Println(sum)

	//continue statement

	// for i := 0; i <= n; i++ {
	// 	if i == 5 {
	// 		continue
	// 	}

	// 	fmt.Println(i)
	// }

	//infinity and break statement loops

	i := 1
	for {
		fmt.Println(i)
		if i == 10 {
			break
		}
		i++
	}

}
