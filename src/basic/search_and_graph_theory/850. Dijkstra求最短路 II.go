package main

import (
	"container/heap"
	"fmt"
)

type Heap [][2]int

func (h Heap) Len() int { return len(h) }

func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h Heap) Less(i, j int) bool {
	if h[i][0] != h[j][0] {
		return h[i][0] < h[j][0]
	}
	return h[i][1] < h[j][1]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *Heap) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

func main() {
	const (
		N   = 150010
		Inf = 0x3f3f3f3f
	)
	var (
		g, e, ne, w, dist  [N]int
		st                 [N]bool
		n, m, x, y, z, idx int
		h                  = make(Heap, 0)
	)

	init := func(n int) {
		for i := 0; i <= n; i++ {
			g[i] = -1
			dist[i] = Inf
		}
		idx = 0
	}

	add := func(a, b, c int) {
		e[idx] = b
		ne[idx] = g[a]
		w[idx] = c
		g[a] = idx
		idx++
	}

	fmt.Scanf("%d %d", &n, &m)
	init(n)

	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d %d", &x, &y, &z)
		add(x, y, z)
	}

	dijkstra := func(s int) {
		heap.Init(&h)

		dist[s] = 0
		heap.Push(&h, [2]int{0, s})
		for h.Len() > 0 {
			u := heap.Pop(&h).([2]int)
			if st[u[1]] {
				continue
			}
			st[u[1]] = true

			for i := g[u[1]]; i != -1; i = ne[i] {
				j := e[i]
				if dist[j] > dist[u[1]]+w[i] {
					dist[j] = dist[u[1]] + w[i]
					heap.Push(&h, [2]int{dist[j], j})
				}
			}
		}
	}

	dijkstra(1)
	if dist[n] == Inf {
		fmt.Println(-1)
	} else {
		fmt.Println(dist[n])
	}
}
