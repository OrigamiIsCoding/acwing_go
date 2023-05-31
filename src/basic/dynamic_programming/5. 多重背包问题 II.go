package main

import "fmt"

func main() {
	const N = 2010
	var (
		f             [N]int
		n, m, s, v, w int
	)

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	fmt.Scanf("%d %d", &n, &m)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d", &v, &w, &s)

		base := 1
		for s >= base {
			for j, vv, ww := m, v*base, w*base; j >= vv; j-- {
				f[j] = max(f[j], f[j-vv]+ww)
			}
			s -= base
			base <<= 1
		}
		if s > 0 {
			for j, vv, ww := m, v*s, w*s; j >= vv; j-- {
				f[j] = max(f[j], f[j-vv]+ww)
			}
		}
	}
	fmt.Println(f[m])
}
