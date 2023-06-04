package main

import "fmt"

func main() {
	const N = 110
	var (
		f                   [N][N]int
		n, v, m, v0, m0, w0 int
	)

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	fmt.Scan(&n, &v, &m)
	for i := 0; i < n; i++ {
		fmt.Scan(&v0, &m0, &w0)
		for j := v; j >= v0; j-- {
			for k := m; k >= m0; k-- {
				f[j][k] = max(f[j][k], f[j-v0][k-m0]+w0)
			}
		}
	}
	fmt.Println(f[v][m])
}
