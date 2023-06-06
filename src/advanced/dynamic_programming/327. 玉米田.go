package main

import "fmt"

func main() {
	const (
		N   = 14
		M   = 1 << N
		Mod = 100000000
	)
	var (
		f       [N][M]int
		state   []int
		st      [N]int
		ne      [M][]int
		m, n, t int
	)

	check := func(x int) bool {
		return ((x>>1)&x) == 0 && ((x<<1)&x) == 0
	}

	fmt.Scan(&m, &n)
	for i := 1; i <= m; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&t)
			if t == 0 {
				st[i] |= (1 << j)
			}
		}
	}

	for i := 0; i < 1<<n; i++ {
		if check(i) {
			state = append(state, i)
		}
	}

	for _, a := range state {
		for _, b := range state {
			if (a & b) == 0 {
				ne[a] = append(ne[a], b)
			}
		}
	}

	f[0][0] = 1
	for i := 1; i <= m+1; i++ {
		for _, a := range state {
			for _, b := range ne[a] {
				if (st[i] & a) == 0 {
					f[i][a] = (f[i][a] + f[i-1][b]) % Mod
				}
			}
		}
	}
	fmt.Println(f[m+1][0])
}
