package main

import "fmt"

func main() {
	const (
		N   = 1e5 + 10
		Inf = 0x3f3f3f3f
	)
	var (
		g, e, ne, w, dist, q       [N]int
		in                         [N]bool
		n, m, x, y, z, idx, hh, tt int
	)

	init := func(n int) {
		for i := 0; i <= n; i++ {
			g[i] = -1
			dist[i] = Inf
		}
	}

	add := func(a, b, c int) {
		e[idx], ne[idx], w[idx], g[a] = b, g[a], c, idx
		idx++
	}

	fmt.Scanf("%d %d", &n, &m)
	init(n)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d %d", &x, &y, &z)
		add(x, y, z)
	}

	spfa := func(s int) {
		dist[s] = 0
		q[tt] = s
		tt++
		in[s] = true
		for tt != hh {
			u := q[hh]
			hh++
			if hh == N {
				hh = 0
			}
			in[u] = false

			for i := g[u]; i != -1; i = ne[i] {
				j := e[i]
				if dist[j] > dist[u]+w[i] {
					dist[j] = dist[u] + w[i]
					if !in[j] {
						in[j] = true
						q[tt] = j
						tt++
						if tt == N {
							tt = 0
						}
					}
				}
			}
		}
	}

	spfa(1)

	if dist[n] == Inf {
		fmt.Println("impossible")
	} else {
		fmt.Println(dist[n])
	}
}
