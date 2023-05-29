package main

import "fmt"

func main() {
	const (
		N   = 510
		Inf = 0x3f3f3f3f
	)
	var (
		g             [N][N]int
		dist          [N]int
		st            [N]bool
		n, m, x, y, z int
	)
	init := func(n int) {
		for i := 1; i <= n; i++ {
			dist[i] = Inf
			for j := 1; j <= n; j++ {
				g[i][j] = Inf
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

	fmt.Scanf("%d %d", &n, &m)
	init(n)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d %d", &x, &y, &z)
		g[x][y] = min(g[x][y], z)
	}

	dijkstra := func(s int) {
		dist[s] = 0
		for i := 0; i < n; i++ {
			u := -1
			for j := 1; j <= n; j++ {
				if !st[j] && (u == -1 || dist[u] > dist[j]) {
					u = j
				}
			}
			st[u] = true

			for j := 1; j <= n; j++ {
				dist[j] = min(dist[j], dist[u]+g[u][j])
			}
		}
	}

	dijkstra(1)
	if dist[n] == Inf {
		fmt.Println(-1)
	} else {
		fmt.Println(dist[n])
	}
}
