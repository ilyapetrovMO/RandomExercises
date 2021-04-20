package main

import "fmt"

func main() {
	var x, p, y float32
	fmt.Scan(&x, &p, &y)

	var yr int
	for yr = 0; ; yr++ {
		if x >= y {
			break
		}

		x += x / 100 * p
	}

	fmt.Println(yr)
}
