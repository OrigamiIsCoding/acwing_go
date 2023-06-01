package main

import "fmt"

func main() {
	const N = 1010
	var (
		f    [N][N]int // 表示 0~i 和 0~j 相同子序列的最长长度
		a, b string
		n, m int
	)

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	fmt.Scanf("%d %d\n%s\n%s", &n, &m, &a, &b)
	a = " " + a
	b = " " + b
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i] == b[j] {
				f[i][j] = max(f[i][j], f[i-1][j-1]+1)
			} else {
				f[i][j] = max(f[i-1][j-1], max(f[i-1][j], f[i][j-1]))
			}
		}
	}
	fmt.Println(f[n][m])
}
