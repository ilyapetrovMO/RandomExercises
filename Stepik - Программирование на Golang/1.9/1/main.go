package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	if a%10 == a%100/10 || a%10 == a/100 || a%100/10 == a/100 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
