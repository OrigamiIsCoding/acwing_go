package main

import "fmt"

const (
	N = 1e5 + 10
	M = N * 2
)

var (
	g, st           [N]int
	e, ne           [M]int
	n, m, u, v, idx int
)

func initG(n int) {
	for i := 0; i <= n; i++ {
		g[i] = -1
	}
}

func add(a, b int) {
	e[idx] = b
	ne[idx] = g[a]
	g[a] = idx
	idx++
}

func dfs(u, fa, c int) bool {
	if st[u] != 0 {
		return st[u] == c
	}
	st[u] = c
	for i := g[u]; i != -1; i = ne[i] {
		j := e[i]
		if j != fa {
			if !dfs(j, u, 3-c) {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Scanf("%d %d", &n, &m)

	initG(n)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d", &u, &v)
		add(u, v)
		add(v, u)
	}

	for i := 1; i <= n; i++ {
		if st[i] == 0 {
			if !dfs(i, -1, 1) {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}
