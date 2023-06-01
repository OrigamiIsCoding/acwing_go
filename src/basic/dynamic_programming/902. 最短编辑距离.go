package main

import "fmt"

func main() {
	const N = 1010
	var (
		f    [N][N]int // 将 0~i 变成 0~j 的最小编辑距离
		a, b string
		n, m int
	)

	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}

	fmt.Scanf("%d\n%s\n%d\n%s", &n, &a, &m, &b)
	a = " " + a
	b = " " + b

	for i := 1; i <= n; i++ {
		f[i][0] = i
	}
	for i := 1; i <= m; i++ {
		f[0][i] = i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i] == b[j] {
				f[i][j] = f[i-1][j-1]
			} else {
				f[i][j] = min(f[i-1][j-1], min(f[i-1][j], f[i][j-1])) + 1
			}
		}
	}
	fmt.Println(f[n][m])
}
