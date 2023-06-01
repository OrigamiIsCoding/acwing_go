package main

import "fmt"

func main() {
	const (
		N = 1010
		M = 20
	)
	var (
		q       [N]string
		s       string
		f       [M][M]int
		n, m, d int
	)

	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}

	g := func(a, b string) int {
		n, m := len(a), len(b)
		a = " " + a
		b = " " + b

		for i := 1; i <= n; i++ {
			f[i][0] = i
		}
		for i := 1; i <= m; i++ {
			f[0][i] = i
		}
		for i := 1; i <= n; i++ {
			for j := 1; j <= m; j++ {
				if a[i] == b[j] {
					f[i][j] = f[i-1][j-1]
				} else {
					f[i][j] = min(f[i-1][j-1], min(f[i-1][j], f[i][j-1])) + 1
				}
			}
		}
		return f[n][m]
	}

	find := func(s string, d int) int {
		count := 0
		for i := 0; i < n; i++ {
			if g(s, q[i]) <= d {
				count++
			}
		}
		return count
	}

	fmt.Scanf("%d %d", &n, &m)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &q[i])
	}
	for i := 0; i < m; i++ {
		fmt.Scanf("%s %d", &s, &d)

		fmt.Println(find(s, d))
	}
}
