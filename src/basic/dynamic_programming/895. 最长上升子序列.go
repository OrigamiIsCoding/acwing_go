package main

import "fmt"

func main() {
	const N = 1010
	var (
		q, f [N]int
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
	ans := 0
	for i := 0; i < n; i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if q[j] < q[i] {
				f[i] = max(f[i], f[j]+1)
			}
		}
		ans = max(f[i], ans)
	}
	fmt.Println(ans)
}
