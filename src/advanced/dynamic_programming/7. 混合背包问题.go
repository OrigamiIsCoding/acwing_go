package main

import "fmt"

func main() {
	const N = 1010
	var (
		f                [N]int
		n, v, v0, w0, s0 int
	)

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	update := func(i, v, w int) {
		f[i] = max(f[i], f[i-v]+w)
	}

	fmt.Scan(&n, &v)
	for i := 0; i < n; i++ {
		fmt.Scan(&v0, &w0, &s0)
		if s0 == -1 {
			for j := v; j >= v0; j-- {
				update(j, v0, w0)
			}
		} else if s0 == 0 {
			for j := v0; j <= v; j++ {
				update(j, v0, w0)
			}
		} else {
			for base := 1; s0 >= base; s0, base = s0-base, base*2 {
				for j := v; j >= v0*base; j-- {
					update(j, v0*base, w0*base)
				}
			}
			if s0 > 0 {
				for j := v; j >= v0*s0; j-- {
					update(j, v0*s0, w0*s0)
				}
			}
		}
	}
	fmt.Println(f[v])
}
