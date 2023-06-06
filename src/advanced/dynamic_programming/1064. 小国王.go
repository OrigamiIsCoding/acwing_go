package main

import "fmt"

func main() {
	const (
		N = 12
		M = 1 << N
		K = N * N
	)
	var (
		f     [N][M][K]int
		cnt   [M]int
		state []int
		next  [M][]int
		n, k  int
	)

	check := func(x int) bool {
		return ((x>>1)&x) == 0 && ((x<<1)&x) == 0
	}

	count := func(x int) int {
		cnt := 0
		for i := 0; i < 32; i++ {
			if (x >> i & 1) == 1 {
				cnt++
			}
		}
		return cnt
	}

	fmt.Scan(&n, &k)
	for i := 0; i < (1 << n); i++ {
		if check(i) {
			state = append(state, i)
			cnt[i] = count(i)
		}
	}
	for _, a := range state {
		for _, b := range state {
			if (a&b) == 0 && check(a|b) {
				next[a] = append(next[a], b)
			}
		}
	}

	f[0][0][0] = 1
	for i := 1; i <= n+1; i++ {
		for j := 0; j <= k; j++ {
			for _, a := range state {
				for _, b := range next[a] {
					// i -- a
					// j -- b
					if j >= cnt[a] {
						f[i][a][j] += f[i-1][b][j-cnt[a]]
					}
				}
			}
		}
	}
	fmt.Println(f[n+1][0][k])
}
