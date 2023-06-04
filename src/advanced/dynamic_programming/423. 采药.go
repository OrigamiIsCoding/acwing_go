package main

import "fmt"

func main() {
	const N = 1010
	var (
		f          [N]int
		t, m, v, w int
	)

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	fmt.Scan(&t, &m)
	for k := 0; k < m; k++ {
		fmt.Scan(&v, &w)

		for i := t; i >= v; i-- {
			f[i] = max(f[i], f[i-v]+w)
		}
	}
	fmt.Println(f[t])
}
