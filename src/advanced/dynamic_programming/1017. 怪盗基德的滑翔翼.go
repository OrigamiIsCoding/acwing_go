package main

import "fmt"

func main() {
	const N = 110
	var (
		f, h [N]int
		k, n int
	)

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	fmt.Scan(&k)
	for t := 0; t < k; t++ {
		fmt.Scan(&n)
		for i := 0; i < n; i++ {
			fmt.Scan(&h[i])
		}

		ans := 0
		for i := 0; i < n; i++ {
			f[i] = 1
			for j := 0; j < i; j++ {
				if h[i] > h[j] {
					f[i] = max(f[i], f[j]+1)
				}
			}
			ans = max(ans, f[i])
		}
		for i := 0; i < n; i++ {
			f[i] = 1
			for j := 0; j < i; j++ {
				if h[i] < h[j] {
					f[i] = max(f[i], f[j]+1)
				}
			}
			ans = max(ans, f[i])
		}
		fmt.Println(ans)
	}
}
