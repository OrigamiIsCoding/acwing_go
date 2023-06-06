package main

import "fmt"

const N = 210

var (
	p          [N * N]int
	n, m, x, y int
	op         string
)

func get(x, y int) int {
	return x*n + y
}

func find(x int) int {
	if x != p[x] {
		p[x] = find(p[x])
	}
	return p[x]
}

func initP(n int) {
	for i := 0; i <= n; i++ {
		p[i] = i
	}
}

func main() {
	fmt.Scan(&n, &m)
	initP(n * n)
	for i := 1; i <= m; i++ {
		fmt.Scan(&x, &y, &op)
		x--
		y--

		nx, ny := x, y
		if op == "D" {
			nx++
		} else {
			ny++
		}
		pa, pb := find(get(x, y)), find(get(nx, ny))

		if pa != pb {
			p[pa] = pb
		} else {
			fmt.Println(i)
			return
		}
	}
	fmt.Println("draw")
}
