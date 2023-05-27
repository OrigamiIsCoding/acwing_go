package main

import "fmt"

func main() {
	const (
		N = 1010
	)
	var (
		a                          [N][N]int
		x1, x2, y1, y2, n, m, q, c int
	)
	fmt.Scanf("%d %d %d", &n, &m, &q)

	add := func(x1, y1, x2, y2, c int) {
		a[x1][y1] += c
		a[x1][y2+1] -= c
		a[x2+1][y1] -= c
		a[x2+1][y2+1] += c
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Scanf("%d", &c)
			add(i, j, i, j, c)
		}
	}
	for i := 0; i < q; i++ {
		fmt.Scanf("%d %d %d %d %d", &x1, &y1, &x2, &y2, &c)
		add(x1, y1, x2, y2, c)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			a[i][j] += a[i-1][j] + a[i][j-1] - a[i-1][j-1]
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Println()
	}
}
