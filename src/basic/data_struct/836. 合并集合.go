package main

import "fmt"

const N = 1e5 + 10

var (
	p          [N]int
	op         string
	n, m, a, b int
)

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

func main() {
	fmt.Scanf("%d %d", &n, &m)

	initP(n)

	for i := 0; i < m; i++ {
		fmt.Scanf("%s %d %d", &op, &a, &b)
		pa, pb := find(a), find(b)
		if op == "M" {
			if pa != pb {
				p[pa] = pb
			}
		} else {
			if pa == pb {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}
