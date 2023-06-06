package main

import "fmt"

func main() {
	const (
		N = 110
		M = 1 << 10
	)
	var (
		f     [2][M][M]int
		st    [N]int
		state []int
		ne    [M][M][]int
		cnt   [M]int
		n, m  int
		s     string
	)

	check := func(x int) bool {
		return ((x>>1)&x) == 0 && ((x<<1)&x) == 0 && ((x<<2)&x) == 0 && ((x>>2)&x) == 0
	}

	count := func(x int) int {
		cnt := 0
		for i := 0; i < 32; i++ {
			cnt += (x >> i & 1)
		}
		return cnt
	}

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	fmt.Scan(&n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scan(&s)
		for j, ch := range s {
			if ch == 'H' {
				st[i] |= 1 << j
			}
		}
	}
	for i := 0; i < 1<<m; i++ {
		if check(i) {
			state = append(state, i)
			cnt[i] = count(i)
		}
	}
	for _, a := range state {
		for _, b := range state {
			for _, c := range state {
				if (a&b) == 0 && (b&c) == 0 && (a&c) == 0 {
					ne[a][b] = append(ne[a][b], c)
				}
			}
		}
	}

	for i := 1; i <= n+2; i++ {
		for _, a := range state {
			for _, b := range state {
				for _, c := range ne[a][b] {
					// a -- i
					// b -- i - 1
					// c -- i - 2
					if (st[i] & a) == 0 {
						f[i&1][a][b] = max(f[i&1][a][b], f[(i-1)&1][b][c]+cnt[a])
					}
				}
			}
		}
	}
	fmt.Println(f[(n+2)&1][0][0])
}
