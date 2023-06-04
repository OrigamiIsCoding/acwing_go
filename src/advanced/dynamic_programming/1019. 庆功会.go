package main

import (
	"fmt"
)

func main() {
	const N = 6010
	var (
		f             [N]int
		n, m, v, w, s int
	)

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	fmt.Scan(&n, &m)
	for i := 0; i < n; i++ {
		fmt.Scan(&v, &w, &s)
		for t := 0; t < s; t++ {
			for j := m; j >= v; j-- {
				f[j] = max(f[j], f[j-v]+w)
			}
		}
	}
	fmt.Println(f[m])
}
