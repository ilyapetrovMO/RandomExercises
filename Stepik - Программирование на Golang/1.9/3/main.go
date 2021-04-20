package main

import "fmt"

func main() {
	var in int
	fmt.Scan(&in)

	left := in/100000%100000 + in%100000/10000 + in%10000/1000
	right := in%1000/100 + in%100/10 + in%10

	if left == right {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
