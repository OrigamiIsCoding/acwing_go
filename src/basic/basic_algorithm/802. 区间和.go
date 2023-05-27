package main

import (
	"fmt"
	"sort"
)

func main() {
	const (
		N = 1e5 + 10
	)
	var (
		n, m       int
		x, c, l, r [N]int
		alls       = make([]int, 0)
		q          []int
	)
	fmt.Scanf("%d %d", &n, &m)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &x[i], &c[i])
		alls = append(alls, x[i])
	}
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d", &l[i], &r[i])
		alls = append(alls, l[i])
		alls = append(alls, r[i])
	}

	unique := func(a []int) []int {
		sort.Ints(a)
		i := 0
		for j := 1; j < len(a); j++ {
			if a[i] != a[j] {
				i++
				a[i], a[j] = a[j], a[i]
			}
		}
		return a[:i]
	}
	alls = unique(alls)

	find := func(x int) int {
		l, r := 0, len(alls)-1
		for l < r {
			mid := (l + r + 1) / 2
			if alls[mid] > x {
				r = mid - 1
			} else {
				l = mid
			}
		}
		return l + 1
	}

	q = make([]int, len(alls)+1)

	for i := 0; i < n; i++ {
		q[find(x[i])] += c[i]
	}
	for i := 1; i <= len(alls); i++ {
		q[i] += q[i-1]
	}
	for i := 0; i < m; i++ {
		fmt.Printf("%d\n", q[find(r[i])]-q[find(l[i])-1])
	}
}
