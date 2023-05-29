package main

import "fmt"

const (
	N   = 1e5 + 10
	Inf = 0x3f3f3f3f
)

var (
	g, ne, e, dist  [N]int
	n, m, a, b, idx int
)

func initG(n int) {
	idx = 0
	for i := 0; i <= n; i++ {
		g[i] = -1
		dist[i] = Inf
	}
}
func add(a, b int) {
	e[idx] = b
	ne[idx] = g[a]
	g[a] = idx
	idx++
}

func main() {
	fmt.Scanf("%d %d", &n, &m)
	initG(n)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d", &a, &b)
		add(a, b)
	}

	q := make([]int, 0)
	dist[1] = 0
	q = append(q, 1)
	for len(q) != 0 {
		u := q[0]
		q = q[1:]

		for i := g[u]; i != -1; i = ne[i] {
			j := e[i]
			if dist[j] > dist[u]+1 {
				dist[j] = dist[u] + 1
				q = append(q, j)
			}
		}
	}
	if dist[n] == Inf {
		fmt.Println(-1)
	} else {
		fmt.Println(dist[n])
	}
}
