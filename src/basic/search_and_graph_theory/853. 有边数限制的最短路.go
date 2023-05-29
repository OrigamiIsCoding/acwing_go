package main

import (
	"fmt"
)

type Edge struct {
	a, b, c int
}

func main() {
	const (
		N   = 510
		M   = 10010
		Inf = 0x3f3f3f3f
	)
	var (
		q            [M]Edge
		dist, backup [N]int
		n, m, k      int
	)

	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}

	fmt.Scanf("%d %d %d", &n, &m, &k)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d %d", &q[i].a, &q[i].b, &q[i].c)
	}

	for i := 1; i <= n; i++ {
		dist[i] = Inf
	}
	dist[1] = 0

	for i := 0; i < k; i++ {
		copy(backup[:], dist[:])
		for j := 0; j < m; j++ {
			a, b, c := q[j].a, q[j].b, q[j].c
			dist[b] = min(dist[b], backup[a]+c)
		}
	}

	if dist[n] > Inf/2 {
		fmt.Println("impossible")
	} else {
		fmt.Println(dist[n])
	}
}
