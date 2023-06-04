package main

import "fmt"

func main() {
	const (
		N = 1010
		M = 510
	)
	var (
		f               [N][M]int
		n, m, k, n0, m0 int
	)

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	fmt.Scan(&n, &m, &k)
	for t := 0; t < k; t++ {
		fmt.Scan(&n0, &m0)
		for i := n; i >= n0; i-- {
			for j := m; j >= m0; j-- {
				f[i][j] = max(f[i][j], f[i-n0][j-m0]+1)
			}
		}
	}
	ans := f[n][m-1]
	for i := 0; i <= m; i++ {
		if f[n][i] == ans {
			fmt.Println(ans, m-i)
			break
		}
	}
}
