package main

import "fmt"

func main() {
	const (
		N = 110
		Inf = 0x3f3f3f3f
	)
	var (
		f [N][N]int
		n int
	)

	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}

	fmt.Scan(&n)
	for i := 0; i <= n+1; i++ {
		for j := 0; j <= n+1; j++ {
			f[i][j] = Inf
		}
	}
	f[1][0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Scan(&f[i][j])
			f[i][j] += min(f[i-1][j], f[i][j-1])
		}
	}
	fmt.Println(f[n][n])
}
