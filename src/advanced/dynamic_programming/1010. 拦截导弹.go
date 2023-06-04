package main

import (
	"fmt"
)

func main() {
	const N = 1010
	var (
		f, h [N]int
		n    int
	)

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	for _, err := fmt.Scan(&h[n]); err == nil; _, err = fmt.Scan(&h[n]) {
		n++
	}
	ans := 0
	for i := 0; i < n; i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if h[i] <= h[j] {
				f[i] = max(f[i], f[j]+1)
			}
		}
		ans = max(ans, f[i])
	}
	fmt.Println(ans)

	ans = 0
	for i := 0; i < n; i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if h[i] > h[j] {
				f[i] = max(f[i], f[j]+1)
			}
		}
		ans = max(ans, f[i])
	}
	fmt.Println(ans)
}
