package main

import "fmt"

func main() {
	const (
		N   = 510
		Inf = 0x3f3f3f3f
	)
	var (
		g             [N][N]int
		n, m, u, v, w int
	)

	init := func(n int) {
		for i := 1; i <= n; i++ {
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
		fmt.Scanf("%d %d %d", &u, &v, &w)
		g[u][v] = min(g[u][v], w)
		g[v][u] = min(g[v][u], w)
	}

	prim := func() (int, bool) {
		edges := make([][2]int, 0)
		st := make([]bool, n+1)

		for i := 1; i <= n; i++ {
			if g[1][i] != Inf {
				edges = append(edges, [2]int{i, g[1][i]})
			}
		}
		st[1] = true
		cost := 0

		for i := 0; i < n-1; i++ {
			u, c := -1, -1
			for _, v := range edges {
				v, w := v[0], v[1]
				if !st[v] && (u == -1 || c > w) {
					u = v
					c = w
				}
			}

			if u == -1 {
				return 0, false
			}
			cost += c
			st[u] = true

			for i := 1; i <= n; i++ {
				if g[u][i] != Inf {
					edges = append(edges, [2]int{i, g[u][i]})
				}
			}
		}
		return cost, true
	}

	if v, ok := prim(); ok {
		fmt.Println(v)
	} else {
		fmt.Println("impossible")
	}
}
