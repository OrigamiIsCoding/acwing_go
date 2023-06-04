package main

import "fmt"

func main() {
	const N = 110
	var (
		f, g, t [N]int
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
		fmt.Scan(&t[i])
		f[i] = 1
		for j := 0; j < i; j++ {
			if t[i] > t[j] {
				f[i] = max(f[i], f[j]+1)
			}
		}
	}

	ans := 0
	for i := n - 1; i >= 0; i-- {
		g[i] = 1
		for j := n - 1; j > i; j-- {
			if t[i] > t[j] {
				g[i] = max(g[i], g[j]+1)
			}
		}
		ans = max(ans, f[i]+g[i]-1)
	}
	fmt.Println(n - ans)
}
