package main

import "fmt"

func main() {
	const N = 20010
	var (
		f       [N]int
		n, v, w int
	)

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	fmt.Scan(&v, &n)
	for i := 0; i < n; i++ {
		fmt.Scan(&w)
		for i := v; i >= w; i-- {
			f[i] = max(f[i], f[i-w]+w)
		}
	}
	fmt.Println(v - f[v])
}
