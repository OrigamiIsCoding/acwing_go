package main

import "fmt"

func main() {
	const N = 110
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

		for j := 0; j < s; j++ {
			for k := m; k >= v; k-- {
				f[k] = max(f[k], f[k-v]+w)
			}
		}
	}
	fmt.Println(f[m])
}
