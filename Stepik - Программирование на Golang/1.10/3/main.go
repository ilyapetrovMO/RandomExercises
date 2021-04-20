package main

import "fmt"

func main() {
	var in int
	for {
		fmt.Scan(&in)
		if in > 100 {
			break
		} else if in < 10 {
			continue
		} else {
			fmt.Println(in)
		}
	}
}
