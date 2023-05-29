package main

import "fmt"

func main() {
	const (
		N   = 1e5 + 10
		Inf = 0x3f3f3f3f
	)
	var (
		g, e, ne, w, dist, q, cnt  [N]int
		in                         [N]bool
		n, m, x, y, z, idx, hh, tt int
	)

	init := func(n int) {
		for i := 0; i <= n; i++ {
			g[i] = -1
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

	spfa := func(s int) bool {
		for i := 1; i <= n; i++ {
			in[i] = true
			q[tt] = i
			tt++
		}
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
					cnt[j] = cnt[u] + 1
					if cnt[j] >= n {
						return true
					}
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
		return false
	}

	if spfa(1) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
