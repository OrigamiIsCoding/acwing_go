package main

import "fmt"

func main() {
	const (
		N   = 310
		Inf = 0x3f3f3f3f
	)
	var (
		f [N][N]int // 合并 i~j 需要的代价
		s [N]int
		n int
	)

	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}

	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&s[i])
		s[i] += s[i-1]
	}

	for len := 2; len <= n; len++ {
		for i := 1; i+len-1 <= n; i++ {
			j := i + len - 1

			f[i][j] = Inf
			cost := s[j] - s[i-1]
			for k := i; k+1 <= j; k++ {
				f[i][j] = min(f[i][j], f[i][k]+f[k+1][j]+cost)
			}
		}
	}
	fmt.Println(f[1][n])
}
