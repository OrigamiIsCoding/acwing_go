package main

import "fmt"

const (
	N   = 5e4 + 10
	Mod = 3
)

var (
	p, d           [N]int
	n, k, op, x, y int
)

func initP(n int) {
	for i := 0; i <= n; i++ {
		p[i] = i
		d[i] = 0
	}
}

func find(x int) int {
	if x != p[x] {
		t := find(p[x])
		d[x] = mod(d[x] + d[p[x]])
		p[x] = t
	}
	return p[x]
}

func mod(x int) int {
	return ((x % Mod) + Mod) % Mod
}

func main() {
	fmt.Scanf("%d %d\n", &n, &k)
	initP(n)

	ans := 0
	for i := 0; i < k; i++ {
		fmt.Scanf("%d %d %d\n", &op, &x, &y)

		if x > n || y > n {
			ans++
			continue
		}

		px, py := find(x), find(y)

		// 1 eat 0, 2 eat 1, 0 eat 2
		if op == 1 {
			if px == py {
				// d[px] % Mod = d[py] % Mod
				if mod(d[x]-d[y]) != 0 {
					ans++
				}
			} else {
				// x - d[x] - px - d[px] - ? - py - d[y] - y
				// d[x] + d[px] = d[y]
				p[px] = py
				d[px] = mod(d[y] - d[x])
			}
		} else {
			// x eat y
			if px == py {
				// d[x] + 1 = d[y]
				if mod(d[x]+1-d[y]) != 0 {
					ans++
				}
			} else {
				// x - d[x] - px - d[px] - ? - py - d[y] - y
				// d[x] + d[px]? + 1  = d[y]
				p[px] = py
				d[px] = mod(d[y] - d[x] - 1)
			}
		}
	}
	fmt.Println(ans)
}
