package main

import "fmt"

func main() {
	const N = 1e6 + 10
	var (
		q, a   [N]int
		hh, tt int
		n, k   int
	)
	fmt.Scanf("%d %d", &n, &k)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	run := func(op func(int, int) bool) {
		hh = 0
		tt = -1
		for i := 0; i < n; i++ {
			if tt >= hh && i-q[hh]+1 > k {
				hh++
			}
			for tt >= hh && op(q[tt], i) {
				tt--
			}
			tt++
			q[tt] = i
			if i >= k-1 {
				fmt.Printf("%d ", a[q[hh]])
			}
		}
		fmt.Println()
	}

	run(func(i1, i2 int) bool { return a[i1] >= a[i2] })
	run(func(i1, i2 int) bool { return a[i1] <= a[i2] })
}
