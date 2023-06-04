package main

import "fmt"

func main() {
	const N = 1010
	var (
		f, g, h [N]int
		n       int
	)

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&h[i])

		f[i] = 1
		for j := 0; j < i; j++ {
			if h[i] > h[j] {
				f[i] = max(f[i], f[j]+1)
			}
		}
	}
	ans := 0
	for i := n - 1; i >= 0; i-- {
		g[i] = 1
		for j := n - 1; j > i; j-- {
			if h[i] > h[j] {
				g[i] = max(g[i], g[j]+1)
			}
		}
		ans = max(ans, f[i]+g[i]-1)
	}
	fmt.Println(ans)
}
