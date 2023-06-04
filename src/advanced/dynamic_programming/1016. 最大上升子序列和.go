package main

import "fmt"

func main() {
	const N = 1010
	var (
		f, b [N]int
		n    int
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
		fmt.Scan(&b[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		f[i] = b[i]
		for j := 0; j < i; j++ {
			if b[i] > b[j] {
				f[i] = max(f[i], f[j]+b[i])
			}
		}
		ans = max(ans, f[i])
	}
	fmt.Println(ans)
}
