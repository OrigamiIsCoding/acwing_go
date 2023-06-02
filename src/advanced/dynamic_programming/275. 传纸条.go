package main

import "fmt"

func main() {
	const (
		N = 60
		M = N * 2
	)
	var (
		f    [N][N][M]int
		w    [N][N]int
		n, m int
	)

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	fmt.Scan(&n, &m)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Scan(&w[i][j])
		}
	}

	for k := 2; k <= n+m; k++ {
		for i1 := 1; i1 <= n; i1++ {
			for i2 := 1; i2 <= n; i2++ {
				j1, j2 := k-i1, k-i2
				if j1 >= 1 && j1 <= m && j2 >= 1 && j2 <= m {
					cost := w[i1][j1]
					if i1 != i2 {
						cost += w[i2][j2]
					}
					f[i1][i2][k] = max(
						max(f[i1-1][i2-1][k-1], f[i1][i2][k-1]),
						max(f[i1-1][i2][k-1], f[i1][i2-1][k-1]),
					) + cost
				}
			}
		}
	}
	fmt.Println(f[n][n][n+m])
}
