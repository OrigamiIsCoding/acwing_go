package main

import "fmt"

func main() {
	const N = 200010
	var (
		s, lmin, lmax, y [N]int
		n                int
		ans1, ans2       int64
	)

	lowbit := func(x int) int {
		return x & (-x)
	}

	update := func(i, x int) {
		for j := i; j < N; j += lowbit(j) {
			s[j] += x
		}
	}

	query := func(i int) int {
		sum := 0
		for j := i; j > 0; j -= lowbit(j) {
			sum += s[j]
		}
		return sum
	}

	init := func() {
		for i := 0; i < N; i++ {
			s[i] = 0
		}
	}

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&y[i])
	}
	for i := 0; i < n; i++ {
		lmin[i] = query(y[i] - 1)
		lmax[i] = query(N-1) - query(y[i])
		update(y[i], 1)
	}
	init()
	for i := n - 1; i >= 0; i-- {
		rmin := query(y[i] - 1)
		rmax := query(N-1) - query(y[i])
		ans1 += int64(lmax[i] * rmax)
		ans2 += int64(lmin[i] * rmin)
		update(y[i], 1)
	}
	fmt.Println(ans1, ans2)
}
