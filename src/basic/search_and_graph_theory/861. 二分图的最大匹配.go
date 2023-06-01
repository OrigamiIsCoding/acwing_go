package main

import "fmt"

const (
	N = 1010
	M = 1e5 + 10
)

var (
	g, match             [N]int
	st                   []bool
	e, ne                [M]int
	n1, n2, m, idx, u, v int
)

func initG(n int) {
	for i := 0; i <= n; i++ {
		g[i] = -1
	}
}

func add(a, b int) {
	e[idx], ne[idx], g[a] = b, g[a], idx
	idx++
}

func find(u int) bool {
	st[u] = true
	for i := g[u]; i != -1; i = ne[i] {
		j := e[i]
		if !st[j] {
			st[j] = true
			if match[j] == 0 || find(match[j]) {
				match[j] = u
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Scanf("%d %d %d", &n1, &n2, &m)
	initG(n1 + n2)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d", &u, &v)
		add(u, v+n1)
	}
	count := 0
	for i := 1; i <= n1; i++ {
		st = make([]bool, n1+n2+1)
		if find(i) {
			count++
		}
	}
	fmt.Println(count)
}
