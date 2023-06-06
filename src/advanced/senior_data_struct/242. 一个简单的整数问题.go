package main

import "fmt"

func main() {
	const N = 100010
	var (
		s             [N]int64
		d, a          int64
		l, r, x, n, m int
		op            string
	)

	lowbit := func(x int) int {
		return x & (-x)
	}

	query := func(x int) (sum int64) {
		for i := x; i > 0; i -= lowbit(i) {
			sum += s[i]
		}
		return
	}

	update := func(x int, y int64) {
		for i := x; i <= n; i += lowbit(i) {
			s[i] += y
		}
	}

	fmt.Scan(&n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a)
		update(i, a)
		update(i+1, -a)
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&op)
		if op == "C" {
			fmt.Scan(&l, &r, &d)
			update(l, d)
			update(r+1, -d)
		} else {
			fmt.Scan(&x)
			fmt.Println(query(x))
		}
	}
}
