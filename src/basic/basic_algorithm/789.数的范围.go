package main

import "fmt"

func main() {
	const (
		N = 1e6 + 10
	)
	var (
		n, q int
		a    [N]int
		x    int
	)
	fmt.Scanf("%d %d", &n, &q)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	for i := 0; i < q; i++ {
		fmt.Scanf("%d", &x)

		l, r := 0, n-1
		for l < r {
			mid := (l + r) / 2
			if a[mid] < x {
				l = mid + 1
			} else {
				r = mid
			}
		}
		if a[l] != x {
			fmt.Println("-1 -1")
			continue
		}

		fmt.Printf("%d ", l)
		r = n - 1

		for l < r {
			mid := (l + r + 1) / 2
			if a[mid] > x {
				r = mid - 1
			} else {
				l = mid
			}
		}
		fmt.Printf("%d\n", l)
	}
}
