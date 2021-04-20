package main

import "fmt"

func main() {
	var n, c, d int

	fmt.Scan(&n, &c, &d)

	for i := 1; i <= n; i++ {
		if i%d == 0 {
			continue
		}

		if i%c == 0 {
			fmt.Println(i)
			break
		}
	}
}
