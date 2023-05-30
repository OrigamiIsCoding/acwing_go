package main

import (
	"fmt"
	"sort"
)

const (
	N = 1e5 + 10
	M = 2e5 + 10
)

type Edge struct {
	u, v, w int
}

type Edges [M]Edge

var (
	edges Edges
	p     [N]int
	n, m  int
)

func (e *Edges) Len() int { return m }

func (e *Edges) Less(i, j int) bool {
	if e[i].w != e[j].w {
		return e[i].w < e[j].w
	}
	return i < j
}

func (e *Edges) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func initP(n int) {
	for i := 0; i <= n; i++ {
		p[i] = i
	}
}

func find(x int) int {
	if x != p[x] {
		p[x] = find(p[x])
	}
	return p[x]
}

func kruskal() (int, bool) {
	initP(n)

	for i := 1; i <= n; i++ {
		p[i] = i
	}

	sort.Sort(&edges)

	cost, cnt := 0, 0
	for i := 0; i < m; i++ {
		u, v, w := edges[i].u, edges[i].v, edges[i].w

		pu, pv := find(u), find(v)
		if pu != pv {
			cost += w
			p[pu] = pv
			cnt++
		}
	}
	return cost, cnt == n-1
}

func main() {

	fmt.Scanf("%d %d", &n, &m)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d %d", &edges[i].u, &edges[i].v, &edges[i].w)
	}

	if v, ok := kruskal(); ok {
		fmt.Println(v)
	} else {
		fmt.Println("impossible")
	}
}
