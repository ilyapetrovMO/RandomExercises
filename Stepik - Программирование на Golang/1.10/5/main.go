package main

import "fmt"

func main() {
	var in int
	var arr = make([]int, 0, 5)

	for {
		fmt.Scan(&in)
		if in == 0 {
			break
		}

		arr = append(arr, in)
	}

	var max int
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	count := 0
	for _, v := range arr {
		if v == max {
			count++
		}
	}

	fmt.Println(count)
}
