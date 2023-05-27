package main

import "fmt"

func main() {
	const (
		N = 1e5 + 10
	)
	var (
		n  int
		q  [N]int
		st [N]bool
	)
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	res := 0
	for i, j := 0, 0; j < n; j++ {
		for st[q[j]] {
			st[q[i]] = false
			i++
		}
		st[q[j]] = true
		res = max(res, j-i+1)
	}
	fmt.Println(res)
}
