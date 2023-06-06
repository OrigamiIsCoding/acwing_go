package main

import "fmt"

const (
	N = 2e5 + 10
)

var (
	p             [N]int
	t, n, a, b, e int
	h             map[int]int
	eq, neq       [][2]int
)

func initP(n int) {
	for i := 0; i <= n; i++ {
		p[i] = i
	}
}

func get(x int) int {
	if v, ok := h[x]; ok {
		return v
	} else {
		v := len(h)
		h[x] = v
		return v
	}
}

func find(x int) int {
	if x != p[x] {
		p[x] = find(p[x])
	}
	return p[x]
}

func main() {
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		h = make(map[int]int)
		eq = make([][2]int, 0)
		neq = make([][2]int, 0)
		fmt.Scan(&n)

		for j := 0; j < n; j++ {
			fmt.Scan(&a, &b, &e)
			a, b = get(a), get(b)
			if e == 1 {
				eq = append(eq, [2]int{a, b})
			} else {
				neq = append(neq, [2]int{a, b})
			}
		}

		initP(len(h))

		for _, v := range eq {
			pa, pb := find(v[0]), find(v[1])
			if pa != pb {
				p[pa] = pb
			}
		}

		flag := true
		for _, v := range neq {
			pa, pb := find(v[0]), find(v[1])
			if pa == pb {
				flag = false
				break
			}
		}

		if flag {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
