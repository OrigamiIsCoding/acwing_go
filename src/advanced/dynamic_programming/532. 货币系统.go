package main

import (
	"fmt"
	"sort"
)

func main() {
	const N = 25010
	var (
		f    [N]bool
		w    [N]int
		t, n int
	)

	init := func(n int) {
		for i := 0; i < n; i++ {
			f[i] = false
		}
		f[0] = true
	}

	fmt.Scan(&t)
	for k := 0; k < t; k++ {
		fmt.Scan(&n)
		ans := 0

		for i := 0; i < n; i++ {
			fmt.Scan(&w[i])
		}

		sort.Ints(w[:n])

		init(N)
		for i := 0; i < n; i++ {
			v := w[i]
			if f[v] {
				ans++
			} else {
				for j := v; j < N; j++ {
					f[j] = f[j] || f[j-v]
				}
			}
		}
		fmt.Println(n - ans)
	}
}
