package main

import "fmt"

func main() {
	const (
		N = 1e6 + 10
	)
	var (
		q          [N]int
		n, m, l, r int
	)
	fmt.Scanf("%d %d", &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &q[i])
		q[i] += q[i-1]
	}
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d", &l, &r)
		fmt.Printf("%d\n", q[r]-q[l-1])
	}
}
