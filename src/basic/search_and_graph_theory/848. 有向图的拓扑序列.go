package main

import "fmt"

func main() {
	const N = 1e5 + 10
	var (
		g, e, ne, in, q         [N]int
		n, m, x, y, idx, hh, tt int
	)

	init := func(n int) {
		for i := 0; i <= n; i++ {
			g[i] = -1
		}
		idx = 0
	}

	add := func(x, y int) {
		e[idx] = y
		ne[idx] = g[x]
		g[x] = idx
		idx++
	}

	fmt.Scanf("%d %d", &n, &m)
	init(n)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d", &x, &y)
		add(x, y)
		in[y]++
	}

	hh, tt = 0, -1
	for i := 1; i <= n; i++ {
		if in[i] == 0 {
			tt++
			q[tt] = i
		}
	}
	for tt >= hh {
		u := q[hh]
		hh++
		for i := g[u]; i != -1; i = ne[i] {
			j := e[i]
			in[j]--
			if in[j] == 0 {
				tt++
				q[tt] = j
			}
		}
	}
	if hh == n {
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", q[i])
		}
	} else {
		fmt.Println(-1)
	}
}
