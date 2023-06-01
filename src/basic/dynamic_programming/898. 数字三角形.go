package main

import "fmt"

func main() {
	const (
		N   = 510
		Inf = 1000_000_000
	)
	var (
		f [2][N]int
		n int
	)

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Scanf("%d", &f[i&1][j])
			if j == 1 {
				f[i&1][j] += f[(i-1)&1][j]
			} else if j == i {
				f[i&1][j] += f[(i-1)&1][j-1]
			} else {
				f[i&1][j] += max(f[(i-1)&1][j], f[(i-1)&1][j-1])
			}
		}
	}
	ans := -Inf
	for i := 1; i <= n; i++ {
		ans = max(ans, f[n&1][i])
	}
	fmt.Println(ans)
}
