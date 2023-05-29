package main

import "fmt"

func main() {
	const (
		N   = 210
		Inf = 0x3f3f3f3f
	)
	var (
		dist             [N][N]int
		n, m, k, x, y, z int
	)

	init := func(n int) {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if i != j {
					dist[i][j] = Inf
				}
			}
		}
	}

	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}

	fmt.Scanf("%d %d %d", &n, &m, &k)
	init(n)

	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d %d", &x, &y, &z)
		dist[x][y] = min(dist[x][y], z)
	}

	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	for i := 0; i < k; i++ {
		fmt.Scanf("%d %d", &x, &y)
		if dist[x][y] > Inf/2 {
			fmt.Println("impossible")
		} else {
			fmt.Println(dist[x][y])
		}
	}
}
