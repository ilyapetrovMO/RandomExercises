package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a/10000 != 0 {
		fmt.Println(a / 10000)
	} else if a/1000 != 0 {
		fmt.Println(a / 1000)
	} else if a/100 != 0 {
		fmt.Println(a / 100)
	} else if a/10 != 0 {
		fmt.Println(a / 10)
	} else {
		fmt.Println(a)
	}
}
