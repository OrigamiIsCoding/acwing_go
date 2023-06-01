package main

import "fmt"

func main() {
	const N = 1e5 + 10
	var (
		f, q [N]int
		n    int
	)

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}

	len := 0
	for i := 0; i < n; i++ {
		l, r := 0, len
		for l < r {
			mid := (l + r + 1) / 2
			if f[mid] < q[i] {
				l = mid
			} else {
				r = mid - 1
			}
		}
		len = max(len, l+1)
		f[l+1] = q[i]
	}
	fmt.Println(len)
}
