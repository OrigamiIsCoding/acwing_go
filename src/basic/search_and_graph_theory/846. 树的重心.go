package main

import (
	"fmt"
)

const N = 1e5 + 10

var (
	g                 [N]int
	st                [N]bool
	ne, e             [N * 2]int
	n, a, b, idx, ans int
)

func initG(n int) {
	ans = N
	for i := 0; i <= n; i++ {
		g[i] = -1
	}
	idx = 0
}

func add(a, b int) {
	e[idx] = b
	ne[idx] = g[a]
	g[a] = idx
	idx++
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func dfs(u int) int {
	st[u] = true

	count, res := 1, 0
	for i := g[u]; i != -1; i = ne[i] {
		j := e[i]
		if st[j] {
			continue
		}
		t := dfs(j)
		res = max(res, t)
		count += t
	}
	res = max(res, n-count)
	ans = min(ans, res)
	return count
}

func main() {
	fmt.Scanf("%d", &n)
	initG(n)
	for i := 0; i < n-1; i++ {
		fmt.Scanf("%d %d", &a, &b)
		add(a, b)
		add(b, a)
	}

	dfs(1)
	fmt.Println(ans)
}
