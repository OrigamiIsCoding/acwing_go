package main

import "fmt"

func main() {
	const N = 1010
	var (
		f          [N]int
		n, m, v, w int
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
		fmt.Scanf("%d %d", &v, &w)

		for j := v; j <= m; j++ {
			f[j] = max(f[j], f[j-v]+w)
		}
	}
	fmt.Println(f[m])
}
