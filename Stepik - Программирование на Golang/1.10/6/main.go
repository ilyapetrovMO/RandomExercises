package main

import "fmt"

func main() {
	var c int
	fmt.Scan(&c)
	var arr = make([]int, 0, c)

	var temp int
	for i := 0; i < c; i++ {
		fmt.Scan(&temp)
		arr = append(arr, temp)
	}

	c = 0
	for _, v := range arr {
		var count int
		temp = v
		for count = 1; count < 4; count++ {
			temp = temp / 10
			if temp == 0 {
				break
			}
		}

		if count != 2 {
			continue
		}

		if v%8 == 0 {
			c += v
		}
	}

	fmt.Println(c)
}
