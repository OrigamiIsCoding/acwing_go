package main

import (
	"fmt"
)

func main() {
	const N = 110
	var (
		f, w    [N][N]int
		t, r, c int
	)

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	fmt.Scan(&t)
	for l := 0; l < t; l++ {
		fmt.Scan(&r, &c)
		for i := 1; i <= r; i++ {
			for j := 1; j <= c; j++ {
				fmt.Scan(&w[i][j])
				f[i][j] = w[i][j] + max(f[i-1][j], f[i][j-1])
			}
		}
		fmt.Println(f[r][c])
	}
}
