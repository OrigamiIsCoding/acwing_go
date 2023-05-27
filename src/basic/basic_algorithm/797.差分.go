package main

import "fmt"

func main() {
	const (
		N = 1e5 + 10
	)
	var (
		q                [N]int
		n, m, l, r, c, t int
	)
	fmt.Scanf("%d %d", &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &t)
		q[i] += t
		q[i+1] -= t
	}
	for i := 1; i <= m; i++ {
		fmt.Scanf("%d %d %d", &l, &r, &c)
		q[l] += c
		q[r+1] -= c
	}
	for i := 1; i <= n; i++ {
		q[i] += q[i-1]
		fmt.Printf("%d ", q[i])
	}
}
