package main

import "fmt"

const N = 6010

var (
	g, e, ne, w, in [N]int
	n, k, l, idx    int
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func initG(n int) {
	for i := 1; i <= n; i++ {
		g[i] = -1
	}
}

func add(a, b int) {
	e[idx], ne[idx], g[a] = b, g[a], idx
	idx++
}

// a 不参加, b 参加
func dfs(u int) (a, b int) {
	for i := g[u]; i != -1; i = ne[i] {
		j := e[i]
		ja, jb := dfs(j)
		a += max(ja, jb)
		b += ja
	}
	b += w[u]
	return
}

func main() {
	fmt.Scan(&n)
	initG(n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&w[i])
	}
	for i := 0; i < n-1; i++ {
		fmt.Scan(&l, &k)
		add(k, l)
		in[l]++
	}
	for i := 1; i <= n; i++ {
		if in[i] == 0 {
			fmt.Println(max(dfs(i)))
			break
		}
	}
}
