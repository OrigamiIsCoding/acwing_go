package main

import (
	"fmt"
	"sort"
)

const (
	N   = 1e5 + 10
	INF = int(1e9 + 10)
)

type Ranges [N]Range

var (
	n int
	q Ranges
)

type Range struct {
	l, r int
}

func (a Ranges) Len() int {
	return n
}

func (a Ranges) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Ranges) Less(i, j int) bool {
	if a[i].l != a[j].l {
		return a[i].l < a[j].r
	}
	return a[i].r < a[j].r
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &q[i].l, &q[i].r)
	}

	sort.Sort(q)

	count, r := 0, -INF
	for i := 0; i < n; i++ {
		if q[i].l > r {
			count++
		}
		r = max(r, q[i].r)
	}
	fmt.Println(count)
}
