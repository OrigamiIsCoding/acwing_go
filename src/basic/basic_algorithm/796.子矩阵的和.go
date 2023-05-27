package main

import "fmt"

func main() {
	const (
		N = 1010
	)
	var (
		a                       [N][N]int
		n, m, q, x1, x2, y1, y2 int
	)

	fmt.Scanf("%d %d %d", &n, &m, &q)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Scanf("%d", &a[i][j])
			a[i][j] += a[i-1][j] + a[i][j-1] - a[i-1][j-1]
		}
	}
	for i := 0; i < q; i++ {
		fmt.Scanf("%d %d %d %d", &x1, &y1, &x2, &y2)
		fmt.Printf("%d\n", a[x2][y2]-a[x2][y1-1]-a[x1-1][y2]+a[x1-1][y1-1])
	}
}
