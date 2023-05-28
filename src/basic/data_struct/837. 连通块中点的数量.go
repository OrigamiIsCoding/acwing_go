package main

import "fmt"

const N = 1e5 + 10

var (
	p, cnt     [N]int
	op         string
	n, m, a, b int
)

func initP(n int) {
	for i := 0; i <= n; i++ {
		p[i] = i
		cnt[i] = 1
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
		fmt.Scanf("%s", &op)
		if op == "C" {
			fmt.Scanf("%d %d", &a, &b)
			pa, pb := find(a), find(b)
			if pa != pb {
				p[pa] = pb
				cnt[pb] += cnt[pa]
			}
		} else if op == "Q1" {
			fmt.Scanf("%d %d", &a, &b)
			pa, pb := find(a), find(b)
			if pa == pb {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		} else {
			fmt.Scanf("%d", &a)
			fmt.Println(cnt[find(a)])
		}
	}
}
