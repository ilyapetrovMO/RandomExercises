package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	found := make([]int, 0, 5)
	for div, mod := 1, 10; a/div != 0; div, mod = div*10, mod*10 {
		n1 := a % mod / div

		for div2, mod2 := 1, 10; b/div2 != 0; div2, mod2 = div2*10, mod2*10 {
			n2 := b % mod2 / div2

			if n1 == n2 {
				f := false
				for _, v := range found {
					if n1 == v {
						f = true
						break
					}
				}

				if f {
					continue
				}

				found = append(found, n1)
			}
		}
	}

	for i := len(found) - 1; i >= 0; i-- {
		fmt.Print(found[i], " ")
	}
}
